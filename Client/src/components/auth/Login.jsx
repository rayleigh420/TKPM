import { useContext, useState } from "react"
import AuthContext from "../../context/AuthProvider"
import { useMutation } from "@tanstack/react-query"
import { login } from "../../api/authApi"

const Login = ({ log, setSign, setLog }) => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const { setAuth } = useContext(AuthContext)

    const loginMutate = useMutation({
        mutationFn: ({ email, password }) => login({ email, password }),
        onSuccess: (data) => {
            setEmail('');
            setPassword("")
            console.log(data)
        },
        onError: (err) => {
            console.log("Error: ", err.response.data.error)
        }
    })

    const handleChangeEmail = (e) => {
        setEmail(e.target.value)
    }

    const handleChangePassword = (e) => {
        setPassword(e.target.value)
    }

    const handleSignIn = () => {
        loginMutate.mutate({
            email: email,
            password: password
        })
        // setAuth({
        //     name: "Le Nhat Duy",
        //     email: 'nhatduy0409@gmail.com',
        //     role: email === "admin" ? "admin" : 'user'
        // })
        // setLog(false)
    }

    const handleSignUps = () => {
        setSign(true)
        setLog(false)
    }

    const handleCloseLogIn = () => {
        setLog(false)
    }

    return (
        <>
            <input checked={log} type="checkbox" id="my-modal-4" className="modal-toggle" />
            <label htmlFor="my-modal-4" className="modal cursor-pointer">
                <label className="modal-box relative flex flex-col items-center max-w-md" htmlFor="">
                    <label htmlFor="my-modal-4" className="btn btn-sm btn-circle absolute right-2 top-2" onClick={handleCloseLogIn}>✕</label>
                    <h3 className="text-lg font-bold text-[#ffffff]">Login</h3>
                    <div className="w-full mt-[22px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Email</span>
                            </label>
                            <input value={email} onChange={handleChangeEmail} type="email" placeholder="Enter your email" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[24px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Password</span>
                            </label>
                            <input value={password} onChange={handleChangePassword} type="password" placeholder="Enter your password" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="mt-[20px]">
                        <p className="text-[13px] font-semibold leading-[19.5px]" >Don't have account?
                            <label htmlFor="my-modal-5" >
                                <label className="cursor-pointer text-[#ffffff]" htmlFor="my-modal-4" onClick={handleSignUps}>    Sign Up</label>
                            </label>
                        </p>
                    </div>
                    <button className="btn mt-[20px]" onClick={handleSignIn}>Sign In</button>
                </label>
            </label>
        </>
    )
}

export default Login