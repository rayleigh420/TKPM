import { useState } from "react"

const AddVersionBook = () => {
    const [id, setID] = useState('')
    const [location, setLocation] = useState('')

    const handleChangeID = (e) => {
        setID(e.target.value)
    }

    const handleChangeLocation = (e) => {
        setLocation(e.target.value)
    }

    const handleAddVersion = () => {
        setID('')
        setLocation('')
        console.log(id, location)
    }

    return (
        <>
            {/* The button to open modal */}
            {/* <label htmlFor="my-modal-4" className="btn">open modal</label> */}

            {/* Put this part before </body> tag */}
            <input type="checkbox" id="modal_addVersion" className="modal-toggle" />
            <label htmlFor="modal_addVersion" className="modal cursor-pointer">
                <label className="modal-box relative" htmlFor="">
                    <h3 className="text-lg font-bold text-center">Add new version for book</h3>
                    <div className="w-full mt-[22px]">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">ID version</span>
                            </label>
                            <input value={id} onChange={handleChangeID} type="text" placeholder="Enter new ID" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>

                    <div className="w-full mt-[24px]">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Location</span>
                            </label>
                            <input value={location} onChange={handleChangeLocation} type="text" placeholder="Enter new location " className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>
                    <div className="flex flex-row justify-end mt-[20px]">
                        <button className="btn" onClick={handleAddVersion}>Create</button>
                    </div>
                </label>
            </label>
        </>
    )
}

export default AddVersionBook