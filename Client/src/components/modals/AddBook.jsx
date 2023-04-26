import { useState } from "react"

const AddBook = () => {
    const [img, setImg] = useState(null)
    const [name, setName] = useState('')
    const [author, setAuthor] = useState('')
    const [type, setType] = useState('')
    const [year, setYear] = useState(2023)
    const [producer, setProducer] = useState('')
    const [publishing, setPublishing] = useState('')

    const handleChangeImage = (e) => {
        setImg(URL.createObjectURL(e.target.files[0]))
    }

    const handleChangeAuthor = (e) => {
        setAuthor(e.target.value)
    }

    const handleChangeType = (e) => {
        setType(e.target.value)
    }

    const handleChangeYear = (e) => {
        setYear(e.target.value)
    }

    const handleChangeProducer = (e) => {
        setProducer(e.target.value)
    }

    const handleChangePublishing = (e) => {
        setPublishing(e.target.value)
    }

    const handleChangeName = (e) => {
        setName(e.target.value)
    }

    const handleAddBook = () => {
        console.log(name, author, type, year, producer, publishing)
        setName('')
        setAuthor('')
        setType('')
        setYear('')
        setProducer('')
        setPublishing('')
    }

    return (
        <>
            <input type="checkbox" id="modal_addBook" className="modal-toggle" />
            <label htmlFor="modal_addBook" className="modal cursor-pointer">
                <label className="modal-box relative w-11/12 max-w-3xl" htmlFor="">
                    <h3 className="text-lg font-bold text-center">Add new book</h3>
                    <div className="flex flex-row justify-center mt-[30px]">
                        <div className="avatar">
                            <div className="w-24 rounded-full hover:shadow-md ring hover:ring-[#121314] ring-[#064CF6] ring-offset-base-100 ring-offset-2">
                                <label className="cursor-pointer" htmlFor="file_input">
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
                                <input value={name} onChange={handleChangeName} type="text" placeholder="Enter name of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Author</span>
                                </label>
                                <input value={author} onChange={handleChangeAuthor} type="text" placeholder="Enter author of book " className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-row">
                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Type</span>
                                </label>
                                <select className="select select-bordered w-full max-w-xs" value={type} onChange={handleChangeType}>
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
                                <input value={year} onChange={handleChangeYear} type="number" min="1900" max="2099" step="1" placeholder="Chose year of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>

                    <div className="flex flex-row">
                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Producer</span>
                                </label>
                                <input value={producer} onChange={handleChangeProducer} type="text" placeholder="Enter producer of book " className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Publishing Location</span>
                                </label>
                                <input value={publishing} onChange={handleChangePublishing} type="text" placeholder="Enter publishing location of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-row justify-end mt-[30px]">
                        <label htmlFor="modal_addBook" className="btn" onClick={handleAddBook} >Create</label>
                    </div>
                </label>
            </label>
        </>
    )
}

export default AddBook