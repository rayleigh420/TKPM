import { useState } from "react"

const ProfileForm = () => {
    const [name, setName] = useState('Le Nhat Duy')
    const [email, setEmail] = useState('strip@gmail.com')

    const handleChangeName = (e) => {
        setName(e.target.value)
    }

    const handleChangeEmail = (e) => {
        setEmail(e.target.value)
    }

    const handleSaveProfile = () => {
        console.log(name, email)
    }

    return (

        <>
            <div className="mt-[50px] w-[300px]">
                <div className="w-full mt-[22px] flex justify-center">
                    <div className="form-control w-full max-w-xs">
                        <label className="label">
                            <span className="label-text text-[#ffffff] font-semibold">Name</span>
                        </label>
                        <input value={name} onChange={handleChangeName} type="text" placeholder="Le Nhat Duy" className="input input-bordered w-full max-w-xs" />
                    </div>
                </div>

                <div className="w-full mt-[24px] flex justify-center">
                    <div className="form-control w-full max-w-xs">
                        <label className="label">
                            <span className="label-text text-[#ffffff] font-semibold">Email</span>
                        </label>
                        <input value={email} onChange={handleChangeEmail} type="email" placeholder="strip@gmail.com" className="input input-bordered w-full max-w-xs" />
                    </div>
                </div>
                <button className="lg:ml-[600px] btn w-[125px] mt-[35px] bg-gradient-to-r from-indigo-700 to-blue-700 text-[#ffffff] leading-[24px] hover:from-indigo-600 hover:to-blue-600" onClick={handleSaveProfile}>Save</button>
            </div>
        </>
    )
}

export default ProfileForm