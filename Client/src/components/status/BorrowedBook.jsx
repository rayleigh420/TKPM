import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import { useContext, useState } from "react"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { getBorrowedBook, paidBook, searchBorrowedBook } from "../../api/manageApi"
import { toast } from "react-toastify"
import AuthContext from "../../context/AuthProvider"

const BorrowedBook = () => {
    const [search, setSearch] = useState('')
    const [searchResult, setSearchResult] = useState()

    const { auth } = useContext(AuthContext)

    const queryClient = useQueryClient()

    const { data: books, isLoading, isError } = useQuery({
        queryKey: ['admin', 'borrowed'],
        queryFn: () => getBorrowedBook()
    })

    const paidBookMutate = useMutation({
        mutationFn: (info) => paidBook(info),
        onSuccess: (data) => {
            console.log(data)
            toast.info("Paid Book Success!")
        },
        onError: () => {
            toast.error("Somethings error. Please try again!")
        },
        onSettled: () => {
            queryClient.invalidateQueries({ queryKey: ['admin', 'borrowed'] })
        }
    })

    const searchBorrowedBookMutate = useMutation({
        mutationFn: (search) => searchBorrowedBook(search),
        onSuccess: (data) => {
            if (data?.error) {
                toast.error("Borrowed ID is not exist")
                setSearch('')
                setSearchResult()
            } else {
                setSearchResult(data)
            }
        },
        onError: () => {
            toast.error("Booked ID is not exist")
            setSearch('')
            setSearchResult()
        }
    })

    const handleChangeSeach = (e) => {
        setSearch(e.target.value)
        if (e.target.value == '') {
            setSearchResult()
        }
    }

    const handleSubmitSearch = (e) => {
        e.preventDefault()
        if (search != '') {
            searchBorrowedBookMutate.mutate(search)
        }
    }

    const handleChangePaid = (item) => {
        console.log("Change to borrowed: ", item)
        paidBookMutate.mutate({
            book_hire_id: item,
            token: auth.token
        })
    }

    console.log(books)

    return (
        <>
            <div className="ml-[180px] w-[880px]">
                <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">List Borrowed Book</h1>
                <div className="divider bordered border-[#ffffff] w-[880px]"></div>
                <div className="flex justify-end mb-[20px]">
                    <form className="flex flex-row gap-[20px]" onSubmit={handleSubmitSearch}>
                        <div className="form-control">
                            <input value={search} onChange={handleChangeSeach} type="text" placeholder="Search" className="input input-bordered bg-[#262627] w-[300px]  focus:ease-out" />
                        </div>
                    </form>
                </div>
                <div className="overflow-x-auto overflow-y-auto w-full h-[800px]">
                    <table className="table w-full relative">
                        <thead className="sticky top-0">
                            <tr>
                                <th>Paid</th>
                                <th>User</th>
                                <th>Book</th>
                                <th>From</th>
                                <th>To</th>
                                <th>Status</th>
                            </tr>
                        </thead>
                        <tbody>
                            {/* row 1 */}
                            {
                                !searchResult ? books?.map((item, index) => (
                                    <tr key={item?._id}>
                                        <td>
                                            <label>
                                                <input type="checkbox" className="checkbox" onChange={() => handleChangePaid(item?.book_hire_id)} />
                                            </label>
                                        </td>
                                        <td>
                                            <div className="flex items-center space-x-3">
                                                <div className="avatar">
                                                    <div className="mask mask-squircle w-12 h-12">
                                                        <img src={item?.user.avatar} alt="Avatar Tailwind CSS Component" />
                                                    </div>
                                                </div>
                                                <div>
                                                    {/* <Link to="/book"> */}
                                                    <div className="font-bold">{item?.user.name}</div>
                                                    {/* </Link> */}
                                                </div>
                                            </div>
                                        </td>
                                        <td>
                                            <div className="flex items-center space-x-3">
                                                <div className="avatar">
                                                    <div className="mask mask-squircle w-12 h-12">
                                                        <img src={item?.book.book_img} alt="Avatar Tailwind CSS Component" />
                                                    </div>
                                                </div>
                                                <div>
                                                    <Link to={`/book/${item?.book.book_id}`}>
                                                        <div className="font-bold">{item?.book.name}</div>
                                                        <div className="text-sm opacity-50">{item?.book_detail.book_detail_id}</div>
                                                    </Link>
                                                </div>
                                            </div>
                                        </td>
                                        <td>{item?.date_borrowed}</td>
                                        <td>{item?.date_end}</td>
                                        <td>
                                            <div className="badge bg-red-500 text-[#ffffff] font-semibold capitalize">{item?.book_detail.status}</div>
                                        </td>
                                    </tr>
                                ))
                                    :
                                    (
                                        <tr >
                                            <td>
                                                <label>
                                                    <input type="checkbox" className="checkbox" onChange={() => handleChangePaid(searchResult?.book_hire_id)} />
                                                </label>
                                            </td>
                                            <td>
                                                <div className="flex items-center space-x-3">
                                                    <div className="avatar">
                                                        <div className="mask mask-squircle w-12 h-12">
                                                            <img src={searchResult?.user.avatar} alt="Avatar Tailwind CSS Component" />
                                                        </div>
                                                    </div>
                                                    <div>
                                                        {/* <Link to="/book"> */}
                                                        <div className="font-bold">{searchResult?.user.name}</div>
                                                        {/* </Link> */}
                                                    </div>
                                                </div>
                                            </td>
                                            <td>
                                                <div className="flex items-center space-x-3">
                                                    <div className="avatar">
                                                        <div className="mask mask-squircle w-12 h-12">
                                                            <img src={searchResult?.book.book_img} alt="Avatar Tailwind CSS Component" />
                                                        </div>
                                                    </div>
                                                    <div>
                                                        <Link to={`/book/${searchResult?.book.book_id}`}>
                                                            <div className="font-bold">{searchResult?.book.name}</div>
                                                            <div className="text-sm opacity-50">{searchResult?.book_detail.book_detail_id}</div>
                                                        </Link>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>{searchResult?.date_borrowed}</td>
                                            <td>{searchResult?.date_end}</td>
                                            <td>
                                                <div className="badge bg-red-500 text-[#ffffff] font-semibold capitalize">{searchResult?.book_detail.status}</div>
                                            </td>
                                        </tr>
                                    )
                            }

                        </tbody>
                        {/* foot */}
                        <tfoot className="sticky bottom-0">
                            <tr>
                                <th>Paid</th>
                                <th>User</th>
                                <th>Book</th>
                                <th>From</th>
                                <th>To</th>
                                <th>Status</th>
                            </tr>
                        </tfoot>
                    </table>
                </div>
            </div >
        </>
    )
}

export default BorrowedBook

