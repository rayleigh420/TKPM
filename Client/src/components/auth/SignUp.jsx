import { useMutation } from "@tanstack/react-query"
import { useContext, useState } from "react"
import { signUp } from "../../api/authApi"
import { toast } from "react-toastify"
import AuthContext from "../../context/AuthProvider"

const SignUp = ({ sign, setLog, setSign }) => {
    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [repassword, setRePassword] = useState('')

    const { setAuth } = useContext(AuthContext)

    const signUpMutation = useMutation({
        mutationFn: ({ name, phone, email, password }) => signUp({ name, phone, email, password }),
        onSuccess: (data) => {
            setName('')
            setEmail('')
            setPassword('')
            setRePassword('')
            toast.info("Register Success!")
            setSign(false)
            setAuth({
                ava: data.ava,
                name: data.name,
                email: data.email,
                role: data.role,
                user_id: data.user_id,
                token: data.token
            })
        },
        onError: (err) => {
            toast.warn(err.response.data.error)
        }
    })

    const handleChangeName = (e) => {
        setName(e.target.value)
    }

    const handleChangeEmail = (e) => {
        setEmail(e.target.value)
    }

    const handleChangePassword = (e) => {
        setPassword(e.target.value)
    }

    const handleChangeRePassword = (e) => {
        setRePassword(e.target.value)
    }

    const handleSignUp = () => {
        // console.log(name, email, password, repassword)
        signUpMutation.mutate({
            name: name,
            phone: '9191919',
            email: email,
            password: password
        })
    }

    const handleLogIn = () => {
        setLog(true)
        setSign(false)
        // setName('')
        // setEmail('')
        // setPassword('')
        // setRePassword('')
    }

    const handleCloseSignUp = () => {
        setSign(false)
    }

    return (
        <>
            <input checked={sign} type="checkbox" id="my-modal-5" className="modal-toggle" />
            <label htmlFor="my-modal-5" className="modal cursor-pointer">
                <label className="modal-box relative flex flex-col items-center max-w-md" htmlFor="">
                    <label htmlFor="my-modal-4" className="btn btn-sm btn-circle absolute right-2 top-2" onClick={handleCloseSignUp}>✕</label>
                    <h3 className="text-lg font-bold text-[#ffffff]">Register</h3>
                    <div className="w-full mt-[22px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Name</span>
                            </label>
                            <input value={name} onChange={handleChangeName} type="text" placeholder="Enter your name" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[24px] flex justify-center">
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

                    <div className="w-full mt-[24px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Confirm Password</span>
                            </label>
                            <input value={repassword} onChange={handleChangeRePassword} type="password" placeholder="Re Enter your password" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="mt-[20px]">
                        <p className="text-[13px] font-semibold leading-[19.5px]">Already have an account?
                            <label htmlFor="my-modal-4" >
                                <label className="cursor-pointer text-[#ffffff]" htmlFor="my-modal-5" onClick={handleLogIn}>    Log In</label>
                            </label>
                        </p>
                    </div>
                    <button className="btn mt-[20px]" onClick={handleSignUp}>Sign Up</button>
                </label>
            </label>
        </>
    )
}

export default SignUp