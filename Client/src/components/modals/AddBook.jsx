import { useState } from "react"

const AddBook = () => {
    const [img, setImg] = useState(null)

    const handleChangeImage = (e) => {
        setImg(URL.createObjectURL(e.target.files[0]))
    }

    return (
        <>
            <input type="checkbox" id="modal_addBook" className="modal-toggle" />
            <label htmlFor="modal_addBook" className="modal cursor-pointer">
                <label className="modal-box relative w-11/12 max-w-3xl" htmlFor="">
                    <h3 className="text-lg font-bold text-center">Add new book</h3>
                    <div className="flex flex-row justify-center mt-[30px]">
                        <div className="avatar">
                            <div className="w-24 rounded-full ring ring-[#121314] ring-offset-base-100 ring-offset-2">
                                <label class="cursor-pointer" htmlFor="file_input">
                                    <img src={img || "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/OOjs_UI_icon_userAvatar.svg/1200px-OOjs_UI_icon_userAvatar.svg.png"} />
                                </label>
                                <input className="hidden" id="file_input" type="file" accept=".jpg,.jpeg,.png" onChange={handleChangeImage}></input>
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-row">
                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Name</span>
                                </label>
                                <input type="text" placeholder="Enter name of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Author</span>
                                </label>
                                <input type="text" placeholder="Enter author of book " className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-row">
                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Type</span>
                                </label>
                                <select className="select select-bordered w-full max-w-xs">
                                    <option disabled selected>Enter type of book</option>
                                    <option>Light Novel</option>
                                    <option>Manga</option>
                                </select>
                            </div>
                        </div>

                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Year</span>
                                </label>
                                <input type="number" min="1900" max="2099" step="1" placeholder="Chose year of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>

                    <div className="flex flex-row">
                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Producer</span>
                                </label>
                                <input type="text" placeholder="Enter producer of book " className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Publishing Location</span>
                                </label>
                                <input type="text" placeholder="Enter publishing location of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-row justify-end mt-[30px]">
                        <label htmlFor="modal_addBook" className="btn" >Create</label>
                    </div>
                </label>
            </label>
        </>
    )
}

export default AddBook