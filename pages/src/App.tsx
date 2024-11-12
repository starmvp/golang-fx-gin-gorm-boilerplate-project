import React, { useState, useEffect } from 'react'
import axios, { AxiosError } from 'axios'

// Types for the input and output data
interface RequestResponse {
  input: string
  output: string
}

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false)
  const [userInfo, setUserInfo] = useState<{ name: string; email: string }>(
    localStorage.getItem('user')
      ? JSON.parse(localStorage.getItem('user') as string)
      : { name: '', email: '' }
  )
  const [username, setUsername] = useState<string>('testuser')
  const [password, setPassword] = useState<string>('testpass')
  const [input, setInput] = useState<string>('')
  const [jwtToken, setJwtToken] = useState<string | null>('')
  const [log, setLog] = useState<RequestResponse[]>([])

  useEffect(() => {
    setJwtToken(localStorage.getItem('token') ?? '')
  }, [])

  useEffect(() => {
    if (jwtToken) {
      console.log('jwtToken=', jwtToken)
      localStorage.setItem('token', jwtToken)
      setIsLoggedIn(true)
      ;(async () => {
        const response = await axios.get(
          'http://localhost:27788/api/v1/my/profile',
          {
            headers: {
              Authorization: `Bearer ${jwtToken}`
            }
          }
        )
        setUserInfo(response.data.user.profile)
        localStorage.setItem('user', JSON.stringify(response.data.user.profile))
      })()
    }
  }, [jwtToken])

  // Handle the login
  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:27788/api/v1/login', {
        username,
        password
      })
      const token = response.data.token
      console.log('token=', token)
      if (token) {
        setJwtToken(token)
        alert('Login successful!')
      } else {
        alert('Login failed! no token')
      }
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError
        if (axiosError.response) {
          alert(
            'Request failed: ' + (axiosError.response.data as Error).message
          )
        } else {
          alert('Request failed: ' + axiosError.message)
        }
      } else {
        alert('Request failed: ' + (error as Error).message)
      }
    }
  }

  // Handle sending the input
  const handleRequest = async () => {
    if (!jwtToken) {
      alert('You must log in first!')
      return
    }

    try {
      const response = await axios.post(
        'http://localhost:27788/api/v1/chat',
        { input },
        {
          headers: {
            Authorization: `Bearer ${jwtToken}`
          }
        }
      )

      // Append the input and response to the log
      setLog((prevLog) => [...prevLog, { input, output: response.data.output }])
      setInput('') // Clear the request input
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError
        if (axiosError.response) {
          alert(
            'Request failed: ' + (axiosError.response.data as Error).message
          )
        } else {
          alert('Request failed: ' + axiosError.message)
        }
      } else {
        alert('Request failed: ' + (error as Error).message)
      }
    }
  }

  return (
    <div style={{ height: '100vh', display: 'flex', flexDirection: 'column' }}>
      {isLoggedIn ? (
        <div
          style={{
            marginBottom: '20px',
            borderBottom: '1px solid #ccc',
            padding: '10px'
          }}
        >
          {`${userInfo.name} <${userInfo.email}>`}
        </div>
      ) : (
        <div style={{ marginBottom: '20px', borderBottom: '1px solid #ccc' }}>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            style={{ margin: '5px', padding: '5px' }}
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ margin: '5px', padding: '5px' }}
          />
          <button
            onClick={handleLogin}
            style={{ margin: '5px', padding: '5px' }}
          >
            Login
          </button>
        </div>
      )}

      <div
        style={{
          flexGrow: 1,
          overflowY: 'auto',
          padding: '10px',
          borderBottom: '1px solid #ccc',
          marginTop: '20px'
        }}
      >
        <div style={{ whiteSpace: 'pre-wrap', wordWrap: 'break-word' }}>
          {log.map((entry, index) => (
            <div key={index}>
              <strong>Request:</strong> {entry.input}
              <br />
              <strong>Response:</strong> {entry.output}
              <hr />
            </div>
          ))}
        </div>
      </div>

      <div style={{ marginBottom: '20px' }}>
        <input
          type="text"
          placeholder="Request"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === 'Enter') {
              handleRequest()
            }
          }}
          style={{ margin: '5px', padding: '5px' }}
        />
        <button
          onClick={handleRequest}
          style={{ margin: '5px', padding: '5px' }}
        >
          Send Request
        </button>
      </div>
    </div>
  )
}

export default App
