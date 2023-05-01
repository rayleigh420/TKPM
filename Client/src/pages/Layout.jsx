import { Outlet } from "react-router-dom"
import Footer from "../components/layout/Footer"
import Form from "../components/forms/Form"
import { useMutation } from "@tanstack/react-query"
import { checkToken } from "../api/authApi"
import useLocalStorage from "../hooks/useLocalStorage"
import { useEffect } from "react"

const Layout = () => {
    const [value, setValue] = useLocalStorage("token", '')
    console.log("Token in Layout: ", value)
    const checkTokenMutate = useMutation({
        mutationFn: (token) => checkToken(token),
        onSuccess: (data) => {
            console.log("USer login: ", data)
        }
    })

    useEffect(() => {
        if (value) {
            checkTokenMutate.mutate(value)
        }
    }, [])

    return (
        <>
            <Form />
            <div className="pt-[65px]">
                <Outlet />
            </div>
            <Footer />
        </>
    )
}

export default Layout