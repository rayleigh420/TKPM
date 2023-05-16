import { useState } from "react"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { addVersionBook } from "../../api/manageApi"
import { useParams } from "react-router-dom"
import { toast } from "react-toastify"

const AddVersionBook = () => {
    const { id } = useParams()
    const [location, setLocation] = useState('')

    const queryClient = useQueryClient()

    const addVersionBookMutate = useMutation({
        mutationFn: (info) => addVersionBook(info),
        onSuccess: (data) => {
            console.log(data)
            toast.info("Add more version sucess!")
        },
        onSettled: () => {
            queryClient.invalidateQueries({ queryKey: ['version', id] })
        }
    })

    const handleChangeLocation = (e) => {
        setLocation(e.target.value)
    }

    const handleAddVersion = () => {
        console.log(id, location)
        if (location == '') {
            return;
        }
        addVersionBookMutate.mutate({
            book_id: id,
            location: location
        })
        setLocation('')
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

                    <div className="w-full mt-[24px]">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Location</span>
                            </label>
                            <input required value={location} onChange={handleChangeLocation} type="text" placeholder="Enter new location " className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>
                    <div className="flex flex-row justify-end mt-[20px]">
                        <label htmlFor="modal_addVersion" className="btn" onClick={handleAddVersion}>Create</label>
                    </div>
                </label>
            </label>
        </>
    )
}

export default AddVersionBook