import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { getBorrowedBook } from "../../api/manageApi"

const BorrowedBook = () => {
    const [search, setSearch] = useState('')

    const { data: books, isLoading, isError } = useQuery({
        queryKey: ['admin', 'borrowed'],
        queryFn: () => getBorrowedBook()
    })

    const handleChangeSeach = (e) => {
        setSearch(e.target.value)
    }

    const handleSubmitSearch = (e) => {
        e.preventDefault()
        console.log(search, typeSearch)
    }

    const handleChangePaid = (index) => {
        console.log("Change to borrowed: ", index)
    }

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
                                books?.map((item, index) => (
                                    <tr key={item?._id}>
                                        <td>
                                            <label>
                                                <input type="checkbox" className="checkbox" onChange={() => handleChangePaid(item?._id)} />
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

