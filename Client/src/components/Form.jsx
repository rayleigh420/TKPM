import Login from "./Login"
import SignUp from "./SignUp"
import { useState } from "react"
import Navigation from "./navigation"

const Form = () => {
    const [sign, setSign] = useState(false)
    const [log, setLog] = useState(false)

    // console.log(log, login)

    return (
        <>
            <Login log={log} setSign={setSign} setLog={setLog} />
            <SignUp sign={sign} setLog={setLog} setSign={setSign} />
            <div className="relative">
                <Navigation setLog={setLog} />
            </div>
        </>
    )
}

export default Form