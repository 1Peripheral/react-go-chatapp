import { useState, useEffect } from 'react'

export function ChatCard() {
   const [msg, setMsg] = useState("")
   const [socket, setSocket] = useState(null)

   function sendMsg() {
      socket.send(msg)
   }

   useEffect(() => {
      let conn = new WebSocket("ws://localhost:8080/ws");
      setSocket(conn)
   }, [])

   return (
      <>
         <textarea name="messages_area" id="chat_box" cols="50" rows="4" readOnly></textarea>
         <input type="text" value={msg} onInput={(e) => setMsg(e.target.value)} placeholder='Type something...' />
         <button onClick={() => sendMsg()}>Send</button>
      </>
   )
}