import { useEffect, useState } from 'react'
import './App.css'

interface DashboardData {
  message: string;
  stats: {
    users: number;
    revenue: number;
    growth: string;
  }
}

function App() {
  const [data, setData] = useState<DashboardData | null>(null)

  useEffect(() => {
    // 请求 Go 后端接口
    fetch('http://localhost:8080/api/dashboard')
      .then(res => res.json())
      .then(data => setData(data))
      .catch(err => console.error("Error:", err))
  }, [])

  return (
    <div className="container">
      <h1>🚀 WPS Elite Dashboard (MVP)</h1>
      <p>Backend Status: {data ? "Online ✅" : "Connecting..."}</p>
      
      {data && (
        <div className="stats-grid">
          <div className="card">
            <h3>Total Users</h3>
            <div className="number">{data.stats.users}</div>
          </div>
          <div className="card">
            <h3>Revenue</h3>
            <div className="number">${data.stats.revenue}</div>
          </div>
          <div className="card">
            <h3>Growth</h3>
            <div className="number green">{data.stats.growth}</div>
          </div>
        </div>
      )}
    </div>
  )
}

export default App