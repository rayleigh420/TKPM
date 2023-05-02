import { useContext, useState } from "react"
import AuthContext from "../../context/AuthProvider"
import { useMutation } from "@tanstack/react-query"
import { uploadAva } from "../../api/imgApi"

const ProfileForm = () => {
    const { auth } = useContext(AuthContext)

    const [img, setImg] = useState(null)
    const [imgURL, setImgURL] = useState(auth.ava)
    const [name, setName] = useState(auth.name)
    const [email, setEmail] = useState(auth.email)

    const uploadAvaMutation = useMutation({
        mutationFn: (imgURL) => uploadAva(imgURL),
        onSuccess: (data) => {
            console.log(data)
        }
    })

    const handleChangeImage = (e) => {
        setImgURL(URL.createObjectURL(e.target.files[0]))
        setImg(e.target.files[0])
        console.log(e.target.files[0])
    }

    const handleChangeName = (e) => {
        setName(e.target.value)
    }

    const handleChangeEmail = (e) => {
        setEmail(e.target.value)
    }

    const handleSaveProfile = () => {
        const imgForm = new FormData();
        imgForm.append('image', img);
        uploadAvaMutation.mutate(imgForm)

        console.log(name, email)
        setImgURL(null)
        setName('')
        setEmail('')
    }

    return (

        <div className="ml-[180px]">
            <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">Your profile</h1>
            <div className="divider bordered border-[#ffffff] w-[750px]"></div>
            <div className="mt-[50px] w-[300px]">
                <div className="flex flex-row justify-center mt-[30px]">
                    <div className="avatar">
                        <div className="w-24 rounded-full hover:shadow-md ring ring-[#121314] hover:ring-[#eeeeee] ring-offset-base-100 ring-offset-2">
                            <label className="cursor-pointer" htmlFor="ava_input">
                                <img src={imgURL || "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/OOjs_UI_icon_userAvatar.svg/1200px-OOjs_UI_icon_userAvatar.svg.png"} className={!img && "bg-[#661AE6]"} />
                            </label>
                            <input className="hidden" id="ava_input" type="file" accept=".jpg,.jpeg,.png" onChange={handleChangeImage}></input>
                        </div>
                    </div>
                </div>

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
        </div>
    )
}

export default ProfileForm