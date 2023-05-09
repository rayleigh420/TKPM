import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import AddBook from "../modals/AddBook"
import EditBook from "../modals/EditBook"
import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { getPaidBook } from "../../api/manageApi"

const PaidBook = () => {
    const [search, setSearch] = useState('')

    const { data: books, isLoading, isError } = useQuery({
        queryKey: ['admin', 'paidbook'],
        queryFn: () => getPaidBook()
    })

    const handleChangeSeach = (e) => {
        setSearch(e.target.value)
    }

    const handleSubmitSearch = (e) => {
        e.preventDefault()
        console.log(search, typeSearch)
    }

    return (
        <>
            <div className="ml-[180px] w-[880px]">
                <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">List Paid Book</h1>
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
                                <th>User</th>
                                <th>Book</th>
                                <th>From</th>
                                <th>To</th>
                            </tr>
                        </thead>
                        <tbody>
                            {/* row 1 */}
                            {
                                books.length > 0 && books?.map((item, index) =>
                                (
                                    <tr key={item?.id}>
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
                                                    <Link to="/book">
                                                        <div className="font-bold">{item?.book.book_id}</div>
                                                        <div className="text-sm opacity-50">69696969</div>
                                                    </Link>
                                                </div>
                                            </div>
                                        </td>
                                        <td>{item?.date_borrowed}</td>
                                        <td>{item?.date_return}</td>
                                    </tr>
                                ))
                            }

                        </tbody >
                        {/* foot */}
                        <tfoot tfoot className="sticky bottom-0" >
                            <tr>
                                <th>User</th>
                                <th>Book</th>
                                <th>From</th>
                                <th>To</th>
                            </tr>
                        </tfoot >
                    </table >
                </div >
            </div >
        </>
    )
}

export default PaidBook

