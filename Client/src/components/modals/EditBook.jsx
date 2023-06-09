import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { useEffect, useState } from "react"
import { updateBook } from "../../api/bookApi"
import { toast } from "react-toastify"
import { getListType } from "../../api/typeApi"
import { uploadAva } from "../../api/imgApi"

const EditBook = ({ book }) => {
    console.log(book)
    const [img, setImg] = useState(null)
    const [imgURL, setImgURL] = useState(book?.book_img)
    const [name, setName] = useState(book?.name)
    const [author, setAuthor] = useState(book?.author)
    const [type, setType] = useState(book?.type.typename)
    const [year, setYear] = useState(book?.yearpublished)
    const [page, setPage] = useState(book?.page)
    const [licensed, setLicensed] = useState(book?.license)
    const [producer, setProducer] = useState(book?.publisher)
    const [publishing, setPublishing] = useState(book?.publishing_location)
    const [detail, setDetail] = useState(book?.details)
    const [description, setDescription] = useState(book?.description)

    const queryClient = useQueryClient()

    const { data: types, isLoading, iseError } = useQuery({
        queryKey: ['types'],
        queryFn: () => getListType(),
    })

    const uploadAvaMutation = useMutation({
        mutationFn: (imgForm) => uploadAva(imgForm),
        onSuccess: (data) => {
            console.log(data)
            setImgURL(prev => data)

            updateBookMutate.mutate({
                book_id: book?.book_id,
                info: {
                    "name": name,
                    "publisher": producer,
                    "yearpublished": Number(year),
                    "author": author,
                    "book_img": data,
                    "type_name": type,
                    "page": Number(page),
                    "publishing_location": publishing,
                    "license": licensed,
                    "description": description,
                    "details": detail
                }
            })
        }
    })

    const updateBookMutate = useMutation({
        mutationFn: (data) => updateBook(data),
        onSuccess: (data) => {
            console.log(data)
            toast.info("Update book successed!")
            setImg(null)
            setImgURL('')
            setName('')
            setAuthor('')
            setType('')
            setYear('')
            setPage('')
            setLicensed('')
            setProducer('')
            setPublishing('')
            setDescription('')
        },
        onError: () => {
            toast.error("Something wrong. Please try again")
        },
        onSettled: () => {
            queryClient.invalidateQueries({ queryKey: ['admin', 'books'] })
        }
    })

    useEffect(() => {
        setImg(book?.book_img)
        setName(book?.name)
        setAuthor(book?.author)
        setType(book?.type.typename)
        setYear(book?.yearpublished)
        setPage(book?.page)
        setLicensed(book?.license)
        setProducer(book?.publisher)
        setPublishing(book?.publishing_location)
        setDetail(book?.details)
        setDescription(book?.description)
    }, [book])

    const handleChangeImage = (e) => {
        setImgURL(URL.createObjectURL(e.target.files[0]))
        setImg(e.target.files[0])

        console.log(imgURL)
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

    const handleChangePage = (e) => {
        setPage(e.target.value)
    }

    const handleChangeLicensed = (e) => {
        setLicensed(e.target.value)
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

    const handleChangeDescription = (e) => {
        setDescription(e.target.value)
    }

    const handleChangeDetail = (e) => {
        setDetail(e.target.value)
    }

    const handleUpdateBook = () => {

        if (name == '' || author == '' || type == '' || licensed == '' || producer == '' || publishing == '' || detail == '' || description == '') {
            console.log(name, author, type, year, page, licensed, producer, publishing, detail, description)
            return;
        }

        if (img != null) {
            const imgForm = new FormData();
            imgForm.append('image', img);

            uploadAvaMutation.mutate(imgForm)
        }
        else {
            updateBookMutate.mutate({
                book_id: book?.book_id,
                info: {
                    "name": name,
                    "publisher": producer,
                    "yearpublished": Number(year),
                    "author": author,
                    "book_image": "https://play-lh.googleusercontent.com/wnIIGkeBT-1w5TUAJlnfAbg7Ko8g2Hc4ynTyuEvEGS9gG-OXugAhf1pGNg2BkDC2NA",
                    "type_name": type,
                    "page": Number(page),
                    "publishing_location": publishing,
                    "license": licensed,
                    "description": description,
                    "details": detail
                }
            })
        }
    }

    console.log(imgURL)

    return (
        <>
            <input type="checkbox" id="modal_editBook" className="modal-toggle" />
            <label htmlFor="modal_editBook" className="modal cursor-pointer">
                <label className="modal-box relative w-11/12 max-w-3xl" htmlFor="">
                    <h3 className="text-lg font-bold text-center">Add new book</h3>
                    <div className="flex flex-row justify-center mt-[30px]">
                        <div className="avatar">
                            <div className="w-24 rounded-full hover:shadow-md ring hover:ring-[#121314] ring-[#064CF6] ring-offset-base-100 ring-offset-2">
                                <label className="cursor-pointer" htmlFor="file_input">
                                    <img src={imgURL || "https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/OOjs_UI_icon_userAvatar.svg/1200px-OOjs_UI_icon_userAvatar.svg.png"} />
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
                                <input required value={name} onChange={handleChangeName} type="text" placeholder="Enter name of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Author</span>
                                </label>
                                <input required value={author} onChange={handleChangeAuthor} type="text" placeholder="Enter author of book " className="input input-bordered w-full max-w-xs" />
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
                                    {
                                        types && types?.map(item => (
                                            <option key={item?.typename} value={item?.typename}>{item?.typename}</option>
                                        ))
                                    }
                                </select>
                            </div>
                        </div>

                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Year</span>
                                </label>
                                <input required value={year} onChange={handleChangeYear} type="number" min="1900" max="2099" step="1" placeholder="Chose year of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>

                    <div className="flex flex-row">
                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Producer</span>
                                </label>
                                <input required value={producer} onChange={handleChangeProducer} type="text" placeholder="Enter producer of book " className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Publishing Location</span>
                                </label>
                                <input required value={publishing} onChange={handleChangePublishing} type="text" placeholder="Enter publishing location of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                    </div>

                    <div className="flex flex-row">

                        <div className="w-full mt-[24px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Pages</span>
                                </label>
                                <input required value={page} onChange={handleChangePage} type="number" min="1" max="3000" step="1" placeholder="Enter page of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>
                        <div className="w-full mt-[22px]">
                            <div className="form-control w-full max-w-xs">
                                <label className="label">
                                    <span className="label-text text-[#ffffff]">Licensed</span>
                                </label>
                                <input required value={licensed} onChange={handleChangeLicensed} type="text" placeholder="Enter licensed of book" className="input input-bordered w-full max-w-xs" />
                            </div>
                        </div>

                    </div>

                    <div className="w-full mt-[22px]">
                        <div className="form-control w-full">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Detail</span>
                            </label>
                            <textarea value={detail} onChange={handleChangeDetail} className="textarea textarea-bordered" placeholder="Enter description of book"></textarea>
                        </div>
                    </div>
                    <div className="w-full mt-[22px]">
                        <div className="form-control w-full">
                            <label className="label">
                                <span className="label-text text-[#ffffff]">Description</span>
                            </label>
                            <textarea value={description} onChange={handleChangeDescription} className="textarea textarea-bordered" placeholder="Enter description of book"></textarea>
                        </div>
                    </div>
                    <div className="flex flex-row justify-end mt-[30px]">
                        <label htmlFor="modal_editBook" className="btn" onClick={handleUpdateBook} >Update</label>
                    </div>
                </label>
            </label>
        </>
    )
}

export default EditBook