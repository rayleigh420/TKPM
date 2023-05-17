import { useMutation } from "@tanstack/react-query"
import { useState } from "react"
import { useParams } from "react-router-dom"
import { rentBookAdmin } from "../../api/bookApi"
import { toast } from "react-toastify"

const DirectHire = () => {
    const [email, setEmail] = useState('')
    const { id } = useParams()

    const rentBookAdminMutate = useMutation({
        mutationFn: (info) => rentBookAdmin(info),
        onSuccess: () => {
            toast.info('Rent book successed!')
        },
        onError: (err) => {
            console.log(err)
            toast.error("Something wrong. Please try again!")
        }
    })

    const handleChangeEmail = (e) => {
        setEmail(e.target.value)
    }

    const handleDirectHire = () => {
        console.log(email, id)
        rentBookAdminMutate.mutate({
            book_id: id,
            email: email
        })
        setEmail('')
    }

    return (
        <>
            <input type="checkbox" id="directhire" className="modal-toggle" />
            <label htmlFor="directhire" className="modal cursor-pointer">
                <label className="modal-box relative" htmlFor="">
                    <h3 className="text-lg font-bold text-center">Direct hire for user</h3>

                    <div className="w-full mt-[24px]">
                        <div className="form-control w-full max-w-xs">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Email</span>
                            </label>
                            <input required value={email} onChange={handleChangeEmail} type="email" placeholder="Enter email" className="input input-bordered w-full max-w-xs" />
                        </div>
                    </div>
                    <div className="flex flex-row justify-end mt-[20px]">
                        <label htmlFor="directhire" className="btn" onClick={handleDirectHire}>Hire</label>
                    </div>
                </label>
            </label>
        </>
    )
}

export default DirectHire