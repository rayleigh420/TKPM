import { useContext, useEffect, useState } from "react"
import AuthContext from "../context/AuthProvider"
import { toast } from "react-toastify";
import useLocalStorage from "./useLocalStorage";

const useLogout = (log = false) => {
    const { setAuth } = useContext(AuthContext);
    const [value, setValue] = useLocalStorage('token', '')

    const [logout, setLogout] = useState(false);
    console.log("logout", logout)

    useEffect(() => {
        if (logout) {
            setAuth({})
            setValue("")
            toast.info("Logout success!")
            setLogout(false)
        }
    }, [logout])

    return setLogout
}

export default useLogout