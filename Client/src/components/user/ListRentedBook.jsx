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
                            data.map(item => (
                                <tr>
                                    <td>
                                        <div className="flex items-center space-x-3">
                                            <div className="avatar">
                                                <div className="mask mask-squircle w-12 h-12">
                                                    <img src={item.src} alt="Avatar Tailwind CSS Component" />
                                                </div>
                                            </div>
                                            <div>
                                                <Link to="/book">
                                                    <div className="font-bold">{item.name}</div>
                                                </Link>
                                            </div>
                                        </div>
                                    </td>
                                    <td>69696969</td>
                                    <td>12/03/2023</td>
                                    <td>12/05/2023</td>
                                    <td>
                                        <div className="badge bg-red-500 text-[#ffffff] font-semibold">Waiting</div>
                                    </td>
                                </tr>
                            ))
                        }
                        {
                            data.map(item => (
                                <tr>
                                    <td>
                                        <div className="flex items-center space-x-3">
                                            <div className="avatar">
                                                <div className="mask mask-squircle w-12 h-12">
                                                    <img src={item.src} alt="Avatar Tailwind CSS Component" />
                                                </div>
                                            </div>
                                            <div>
                                                <Link to="/book">
                                                    <div className="font-bold">{item.name}</div>
                                                </Link>
                                            </div>
                                        </div>
                                    </td>
                                    <td>69696969</td>
                                    <td>12/03/2023</td>
                                    <td>12/05/2023</td>
                                    <td>
                                        <div className="badge bg-red-500 text-[#ffffff] font-semibold">Waiting</div>
                                    </td>
                                </tr>
                            ))
                        }
                    </tbody>
                    {/* foot */}
                    <tfoot className="sticky bottom-0">
                        <tr>
                            <th>Name Book</th>
                            <th>ID Book</th>
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