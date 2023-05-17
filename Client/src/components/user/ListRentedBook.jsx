import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import { useContext } from "react"
import { useQuery } from "@tanstack/react-query"
import { getListRentedBook } from "../../api/manageApi"
import AuthContext from "../../context/AuthProvider"

const ListRentedBook = () => {
    const { auth } = useContext(AuthContext)

    const { data: books, isLoading, isError } = useQuery({
        queryKey: ['user', 'history'],
        queryFn: () => getListRentedBook(auth?.user_id)
    })

    console.log(books)

    return (
        <div className="ml-[180px] w-[880px]" >
            <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">List rented book</h1>
            <div className="divider bordered border-[#ffffff] w-[880px]"></div>
            <div className="overflow-x-auto overflow-y-auto w-full h-[800px]">
                <table className="table w-full relative">
                    <thead className="sticky top-0">
                        <tr>
                            <th>Name Book</th>
                            <th>ID Book</th>
                            <th>From</th>
                            <th>To</th>
                            <th>Status</th>
                        </tr>
                    </thead>
                    <tbody className="mt-[30px]">
                        {/* row 1 */}
                        {
                            books && books?.map(item => (
                                <tr key={item?.history_id}>
                                    <td>
                                        <div className="flex items-center space-x-3">
                                            <div className="avatar">
                                                <div className="mask mask-squircle w-12 h-12">
                                                    <img src={item?.book.book_img} alt="Avatar Tailwind CSS Component" />
                                                </div>
                                            </div>
                                            <div>
                                                <Link to={`/book/${item.book.book_id}`}>
                                                    <div className="font-bold">{item?.book.name}</div>
                                                    <div className="text-sm opacity-50">Code: {item?.book_rent_id || item?.book_hire_id || item?.history_id}</div>
                                                </Link>
                                            </div>
                                        </div>
                                    </td>
                                    <td>{item?.book_detail_id}</td>
                                    <td>{item?.date_borrowed}</td>
                                    <td>{item?.date_return || item?.reserve_date}</td>
                                    <td>
                                        {
                                            item?.status == 'returned'
                                                ?
                                                (
                                                    <div className="badge bg-green-500 text-[#ffffff] font-semibold">Returned</div>
                                                )
                                                :
                                                (
                                                    item?.book_detail.status == 'booked'
                                                        ?
                                                        (
                                                            <div className="badge bg-yellow-500 text-[#ffffff] font-semibold">Booked</div>

                                                        )
                                                        :
                                                        (
                                                            <div className="badge bg-red-500 text-[#ffffff] font-semibold">Borrowed</div>

                                                        )
                                                )
                                        }
                                    </td>
                                </tr>
                            ))
                        }
                    </tbody>
                    {/* foot */}
                    <tfoot className="sticky bottom-0">
                        <tr>
                            <th>Name Book</th>
                            <th>ID</th>
                            <th>From</th>
                            <th>To</th>
                            <th>Status</th>
                        </tr>
                    </tfoot>

                </table>
            </div>
        </div>
    )
}

export default ListRentedBook