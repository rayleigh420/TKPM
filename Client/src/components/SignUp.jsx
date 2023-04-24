const SignUp = ({ sign, setLog, setSign }) => {

    const handleLogIn = () => {
        setLog(true)
        setSign(false)
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
                    <h3 className="text-lg font-bold text-[#ffffff]">Sign Up</h3>
                    <div className="w-full mt-[22px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Name</span>
                            </label>
                            <input type="text" placeholder="Enter your name" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[24px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Email</span>
                            </label>
                            <input type="text" placeholder="Enter your email" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[24px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Password</span>
                            </label>
                            <input type="text" placeholder="Enter your password" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[24px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Confirm Password</span>
                            </label>
                            <input type="text" placeholder="Re Enter your password" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="mt-[20px]">
                        <p className="text-[13px] font-semibold leading-[19.5px]">Already have account?
                            <label htmlFor="my-modal-4" >
                                <label className="cursor-pointer text-[#ffffff]" htmlFor="my-modal-5" onClick={handleLogIn}>    Log In</label>
                            </label>
                        </p>
                    </div>
                    <button className="btn mt-[20px]">Sign Up</button>
                </label>
            </label>
        </>
    )
}

export default SignUp