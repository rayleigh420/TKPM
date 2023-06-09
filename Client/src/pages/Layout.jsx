import { Outlet } from "react-router-dom"
import Footer from "../components/layout/Footer"
import Form from "../components/forms/Form"
import { useMutation } from "@tanstack/react-query"
import { checkToken } from "../api/authApi"
import useLocalStorage from "../hooks/useLocalStorage"
import { useContext, useEffect } from "react"
import AuthContext from "../context/AuthProvider"
import { toast } from "react-toastify"

const Layout = () => {
    const { setAuth } = useContext(AuthContext)
    const [value, setValue] = useLocalStorage("token", '')

    // console.log("Token in Layout: ", value)
    const checkTokenMutate = useMutation({
        mutationFn: (token) => checkToken(token),
        onSuccess: (data) => {
            setAuth({
                ava: data.avatar,
                name: data.name,
                email: data.email,
                role: data.role,
                user_id: data.user_id,
                token: data.token
            })
        },
        onError: () => {
            setAuth({})
            setValue('')
            toast.error("Please login again!")
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