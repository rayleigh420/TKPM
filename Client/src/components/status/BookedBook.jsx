import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import AddBook from "../modals/AddBook"
import EditBook from "../modals/EditBook"
import { useState } from "react"
import { useMutation, useQuery } from "@tanstack/react-query"
import { getBookedBook, searchBookedBook } from "../../api/manageApi"
import { toast } from "react-toastify"

const BookedBook = () => {
    const [search, setSearch] = useState('')
    const [searchResult, setSearchResult] = useState()

    const { data: books, isLoading, isError } = useQuery({
        queryKey: ['admin', 'booked'],
        queryFn: () => getBookedBook(),
    })

    const searchBookedBookMutate = useMutation({
        mutationFn: (search) => searchBookedBook(search),
        onSuccess: (data) => {
            if (data?.error) {
                toast.error("Booked ID is not exist")
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
            searchBookedBookMutate.mutate(search)
        }
    }

    const handleChangeBorrowed = (index) => {
        console.log("Change to borrowed: ", index)
    }

    console.log(books)

    return (
        <>
            <div className="ml-[180px] w-[880px]">
                <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">List Booked Book</h1>
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
                                <th>Borrowed</th>
                                <th>User</th>
                                <th>Book</th>
                                <th>Reverse</th>
                            </tr>
                        </thead>
                        <tbody>
                            {/* row 1 */}
                            {
                                !searchResult ? books?.map((item, index) => (
                                    <tr key={item?._id}>
                                        <td>
                                            <label>
                                                <input type="checkbox" className="checkbox" onChange={() => handleChangeBorrowed(item?.book_rent_id)} />
                                            </label>
                                        </td>
                                        <td>
                                            <div className="flex items-center space-x-3">
                                                <div className="avatar">
                                                    <div className="mask mask-squircle w-12 h-12">
                                                        <img src={item?.user?.avatar} alt="Avatar Tailwind CSS Component" />
                                                    </div>
                                                </div>
                                                <div>
                                                    {/* <Link to="/book"> */}
                                                    <div className="font-bold">{item?.user?.name}</div>
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
                                        <td>{item?.reserve_date}</td>
                                    </tr>
                                ))
                                    :
                                    (
                                        <tr>
                                            <td>
                                                <label>
                                                    <input type="checkbox" className="checkbox" onChange={() => handleChangeBorrowed(searchResult?.book_rent_id)} />
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
                                            <td>{searchResult?.reserve_date}</td>
                                        </tr>
                                    )
                            }
                        </tbody>
                        {/* foot */}
                        <tfoot className="sticky bottom-0">
                            <tr>
                                <th>Borrowed</th>
                                <th>User</th>
                                <th>Book</th>
                                <th>Reverse</th>
                            </tr>
                        </tfoot>
                    </table>
                </div>
            </div >
        </>
    )
}

export default BookedBook
