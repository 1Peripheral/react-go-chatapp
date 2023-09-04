import { useState, useEffect } from 'react'
import { ChatCard } from './components/ChatCard'

function App() {
  const [msg, setMsg] = useState("")
  const [socket, setSocket] = useState(null)
  const [id, setId] = useState("")
  const [msgList, setMsgList] = useState([])

  useEffect(() => {
    setSocket(new WebSocket("ws://localhost:8080/ws"))
    if(socket) {
      socket.onopen = (event) => {
        console.log(event.data)
        setId(event.data)
      }
    }
    return () => ws.close()
  }, [])

  if (socket) {
    socket.onmessage = function (event) {
      setMsgList(currentMsgs => {
        return [...currentMsgs, {
          id: crypto.randomUUID(),
          content: event.data 
        }]
      })
    }
  }

  function sendMsg() {
    socket.send(msg)
    setMsg("")
  }

  return (
    <>
      <ChatCard sendMsg={sendMsg} msg={msg} msgList={msgList} setMsg={setMsg}/>
    </>
  )
}

export default App
