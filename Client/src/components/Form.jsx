import Login from "./auth/Login"
import SignUp from "./auth/SignUp"
import { useState } from "react"
import Navigation from "./layout/Navigation"

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