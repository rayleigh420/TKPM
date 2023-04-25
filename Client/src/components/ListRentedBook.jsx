import { Link } from "react-router-dom"
import { data } from "../pages/home"

const ListRentedBook = () => {
    return (
        <div className="ml-[180px] w-[880px]">
            <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">List rented book</h1>
            <div className="divider bordered border-[#ffffff] w-[850px]"></div>
            <div className="overflow-x-auto overflow-y-auto w-full h-[800px]">
                <table className="table w-full">
                    <thead>
                        <tr>
                            <th>Name Book</th>
                            <th>ID Book</th>
                            <th>From</th>
                            <th>To</th>
                            <th>Status</th>
                        </tr>
                    </thead>
                    <tbody>
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
                    {/* <tfoot>
                        <tr>
                            <th>Name Book</th>
                            <th>ID</th>
                            <th>Date</th>
                            <th>Status</th>
                            <th></th>
                        </tr>
                    </tfoot> */}

                </table>
            </div>
        </div>
    )
}

export default ListRentedBook