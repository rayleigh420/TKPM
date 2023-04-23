const Login = () => {
    return (
        <>
            <input type="checkbox" id="my-modal-4" className="modal-toggle" />
            <label htmlFor="my-modal-4" className="modal cursor-pointer">
                <label className="modal-box relative flex flex-col items-center max-w-md" htmlFor="">
                    <h3 className="text-lg font-bold">Login</h3>
                    <div className="w-full mt-[22px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text">Email</span>
                            </label>
                            <input type="text" placeholder="Enter your email" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[22px] flex justify-center">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text">Password</span>
                            </label>
                            <input type="text" placeholder="Enter your password" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="mt-[20px]">
                        <p className="text-[13px] font-semibold leading-[19.5px]">Don't have account?
                            <a className="link no-underline">   Sign Up</a>
                        </p>
                    </div>
                    <button className="btn mt-[20px]">Sign In</button>
                </label>
            </label>
        </>
    )
}

export default Login