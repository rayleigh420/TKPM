import { useContext, useState } from "react"
import HistoryBook from "../components/admin/HistoryBook"
import AddVersionBook from "../components/modals/AddVersionBook"
import RelatedBook from "../components/list/RelatedBook"
import { Link, useParams } from "react-router-dom"
import { useMutation, useQuery } from "@tanstack/react-query"
import { getDetailBook, rentBook } from "../api/bookApi"
import AuthContext from "../context/AuthProvider"
import LocationVersion from "../components/admin/LocationVersion"
import MessageRent from "../components/modals/MessageRent"
import { toast } from "react-toastify"
import DirectHire from "../components/modals/DirectHire"

//status:

// red: borrowed
// yellow: booked
// green: available

const Book = () => {
    const [tab, setTab] = useState(false)
    const [rent, setRent] = useState()
    const { id } = useParams()
    const { error, setError } = useState('')

    const { auth } = useContext(AuthContext)

    const { data: book, isLoading, isError } = useQuery({
        queryKey: ["book", id],
        queryFn: () => getDetailBook(id)
    })

    const rentBookMutate = useMutation({
        mutationFn: (info) => rentBook(info),
        onSuccess: (data) => {
            setRent(data.book_rent_id)
            console.log("Data: ", data)
        },
        onError: (err) => {
            toast.error(err.response.data.error)
            setRent(err.response.data.error)
        }
    })

    const handleTabAbout = () => {
        setTab(false)
    }

    const handleTabLocation = () => {
        setTab(true)
    }

    const handleRentBook = (item) => {
        // setRent('123123123')
        if (book.amount <= 0) {
            toast.warn("All book is busy. Please wait!")
            return;
        }
        if (auth?.role == 'admin') {
            return;
        }
        rentBookMutate.mutate({
            book_id: id,
            user_id: auth?.user_id
        })
    }
    console.log(book)

    return (
        <>
            {/* {
                rent && 
            } */}

            {auth?.role == 'user' && book?.amount > 0 && <MessageRent rent={rent} />}
            {auth?.role == 'admin' && book?.amount > 0 && <DirectHire />}

            <div className="flex flex-col items-center">
                <div className=" w-[62%] py-[60px] flex flex-row">
                    <div className="min-w-[237px] min-h-[347px] w-[240px] h-[350px] bg-cover hover:border-[1.75px] rounded-[7px] overflow-hidden">
                        <img src={book?.book_img || book?.book_image} className="w-full h-full" />
                    </div>
                    {/* <div className={"w-[240px] h-[350px] bg-cover hover:border-[1.75px] rounded-[7px] " + `bg-[url('${book?.book_img}')]`}></div> */}
                    {/* <div className="w-[240px] h-[350px] bg-cover hover:border-[1.75px] bg-[url('https://fastly.picsum.photos/id/200/200/300.jpg?hmac=XVCLpc2Ddr652IrKMt3L7jISDD4au5O9ZIr3fwBtxo8')] rounded-[7px]">
                    </div> */}
                    <div className="ml-[30px]">
                        {
                            book?.amount > 0
                                ?
                                <div className="badge bg-[#eeeeee] rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Ongoing</div>
                                :
                                <div className="badge bg-red-500 rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Over</div>

                        }
                        <h1 className="mt-[10px] text-[#ffffff] text-[32px] font-bold leading-[36.8px] tracking-[-0.64]">{book?.name}</h1>
                        <p className="mt-[20px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Author: {book?.author}</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Type:
                            <Link to={`/${book?.type_name}`}>
                                {" " + book?.type_name}
                            </Link>
                        </p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Pages: {book?.page}</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Year: {book?.yearpublished}</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Producer: {book?.publisher}</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Publishing Location: {book?.publishing_location}</p>
                        <button disabled={auth?.name == undefined} className="btn w-[150px] mt-[55px] bg-gradient-to-r from-indigo-700 to-blue-700 text-[#ffffff] leading-[24px] hover:from-indigo-600 hover:to-blue-600" onClick={() => handleRentBook(book?.book_name)}>
                            <label htmlFor={auth?.role == 'user' ? "messagerent" : "directhire"} className="bg-transparent border-none cursor-pointer">
                                Rent Book
                            </label>
                        </button>
                    </div>
                </div>
                <div className="tabs w-full tab-bordered flex flex-col items-center">
                    <div className="w-[62%]">
                        <a className={"pb-[40px] text-[#ffffff] text-[20px] font-semibold leading-[24px] tracking-[-0.4px] tab tab-bordered" + (tab ? "tab-active" : "")} onClick={handleTabAbout}>About</a>
                        {
                            auth?.role == 'admin' &&
                            <a className={"pb-[40px] text-[#ffffff] text-[20px] font-semibold leading-[24px] tracking-[-0.4px] tab tab-bordered" + (!tab ? "tab-active" : "")} onClick={handleTabLocation}>Location</a>
                        }
                    </div>
                </div>
                {
                    !tab ? (
                        <div className="w-[62%] mt-[30px]">
                            <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px] mb-[20px]">Information</h4>
                            <div className="overflow-x-auto">
                                <table className="table w-full border-nones">
                                    <thead>
                                        <tr>
                                            <th>Pages</th>
                                            <th>Licensed from</th>
                                            <th>Number of renter</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>{book?.page}</td>
                                            <td>{book?.license}</td>
                                            <td>{book?.borrowed_quantity}</td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                            <div className="divider bordered border-[#ffffff]"></div>
                            <div className="text-[#e0e0e0] leading-[24px] tracking-[-0.32px]">
                                <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px] mb-[25px]">Description</h4>
                                {book?.description}
                                {/* For decades, he had lived as a puppet of the tower.
                                Throughout that time, who knew how many deaths he had encountered?
                                One day, however, his memories returned to him.
                                The First Floor Boss finally awakened */}
                            </div>
                            <div className="divider bordered border-[#ffffff]"></div>
                            <div className="text-[#e0e0e0] leading-[24px] tracking-[-0.32px]">
                                <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px] mb-[25px]">Details</h4>
                                {book?.details}
                                {/* Status in Korean: Completed @ 254 chapters <br />
                                <br />
                                Translator: Yeniverse <br />
                                <br />
                                Editor: Dot, SoSam <br />
                                <br />
                                Release rate: 5 chapters a week with an occasional bonus 6th chapter */}
                            </div>
                            <div className="divider bordered border-[#ffffff]"></div>
                        </div>) :
                        (
                            auth?.role == 'admin' && <LocationVersion id={id} />
                        )
                }
                <div className="w-[62%]">
                    <RelatedBook related={book?.related_books} isLoading={isLoading} />
                </div>
            </div >
            <div className="divider bordered border-[#ffffff]"></div>
        </>
    )
}

export default Book 