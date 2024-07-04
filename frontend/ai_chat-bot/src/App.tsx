import { useState } from 'react'
import './App.css'

function App() {
    const [message, setMessage] = useState("")
    const [array, setArray] = useState([String()].splice(1,))

    const sendMessage = async (variable: String) => {
        try {
            setArray([...array, message])
            const response = await fetch('http://localhost:8080/chat', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    messages: [
                        {
                            role: 'user',
                            content: variable,
                        },
                    ],
                }),
            });

            if (response.ok) {
                const data = await response.json();
                // data.response
                setArray([...array, message, data.response])

            } else {
                console.error('Error sending message:', response.statusText);
            }
        } catch (error) {
            console.error('Error sending message:', error);
        }
    };

    return (
        <>
            <div className='main'>
                {array.map((item, index) => (
                    <div className={`message ${((Number(index)) % 2) == 0 ? 'Right' : 'Left'}`}>{item}</div>
                ))}
                <div className='Left message hidden big'>3213213333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333222.</div>
            </div>




            <div className='input_footer'>
                <input className='input_place' value={message} onChange={(event) => { setMessage(event.target.value) }} />
                <div className='send_button'>
                    <span className='arrow' onClick={() => {
                        if (message !== "") {
                            sendMessage(message)
                            setMessage("")
                        }
                    }}>&#8680;</span>
                </div>
            </div>
        </>
    )
}

export default App
