export function MessageElmement( { user, msg }) {
   return (
      <li className="message" key={msg.id}>
         <p className="user">{ user }</p>
         <p className="content">{ msg.content }</p>
      </li>
   )
} 