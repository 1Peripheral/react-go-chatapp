import { useState, } from 'react'
import { MessageElmement } from './MessageElement'

export function ChatCard({ sendMsg, msg, msgList,setMsg }) {
   return (
      <>
         <div className="message_box">
            <ul className='message_list'>
               {
                  msgList.map(msg_ => {
                     return <MessageElmement msg={msg_} user="Anonymous"/>
                  })
               }
            </ul>  
         </div>
         <input type="text" value={msg} onInput={(e) => setMsg(e.target.value)} placeholder='Type something...' />
         <button onClick={() => sendMsg()}>Send</button>
      </>
   )
}