import { data } from "../pages/home"

const ListRentedBook = () => {
    return (
        <div className="ml-[200px] w-[850px]    ">
            <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">List rented book</h1>
            <div className="divider bordered border-[#ffffff] w-[850px]"></div>
            <div className="overflow-x-auto w-full">
                <table className="table w-full">
                    <thead>
                        <tr>
                            <th>Name Book</th>
                            <th>ID</th>
                            <th>Date</th>
                            <th>Status</th>
                            <th></th>
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
                                                <div className="font-bold">{item.name}</div>
                                            </div>
                                        </div>
                                    </td>
                                    <td>69696969</td>
                                    <td>12/03 - 12/05</td>
                                    <td>Done</td>
                                    <th>
                                        <button className="btn btn-ghost btn-xs">Details</button>
                                    </th>
                                </tr>
                            ))
                        }
                    </tbody>
                    {/* foot */}
                    <tfoot>
                        <tr>
                            <th>Name Book</th>
                            <th>ID</th>
                            <th>Date</th>
                            <th>Status</th>
                            <th></th>
                        </tr>
                    </tfoot>

                </table>
            </div>
        </div>
    )
}

export default ListRentedBook