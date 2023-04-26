import { Link } from "react-router-dom"

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]


// status: 
// booked: yellow
// borrowed: red
// returned: green

const HistoryBook = () => {
    return (
        <>
            {/* <label htmlFor="history_book" className="btn">open modal</label> */}

            {/* Put this part before </body> tag */}
            <input type="checkbox" id="history_book" className="modal-toggle" />
            <div className="modal">
                <div className="modal-box w-11/12 max-w-4xl">
                    <h3 className="font-bold text-lg">Hisotry rented of book</h3>
                    <div className="overflow-y-auto h-[400px] mt-[30px]">
                        <table className="table w-full ">
                            <thead>
                                <tr>
                                    <th>Name User</th>
                                    <th>ID Book</th>
                                    <th>From</th>
                                    <th>To</th>
                                    <th>Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                {/* row 1 */}
                                {
                                    arr.map(item => (
                                        <tr>
                                            <td>
                                                <div className="flex items-center space-x-3">
                                                    <div className="avatar">
                                                        <div className="mask mask-squircle w-12 h-12">
                                                            <img src='https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b' alt="Avatar Tailwind CSS Component" />
                                                        </div>
                                                    </div>
                                                    <div>
                                                        <Link to="/">
                                                            <div className="font-bold">Le Nhat Duy</div>
                                                        </Link>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>69696969</td>
                                            <td>12/03/2023</td>
                                            <td>12/05/2023</td>
                                            <td>
                                                {
                                                    (item % 2 == 0) ?
                                                        (

                                                            <div className="badge bg-red-500 text-[#ffffff] font-semibold">Borrowed</div>
                                                        )
                                                        :
                                                        (
                                                            <div className="badge bg-yellow-500 text-[#ffffff] font-semibold">Booked</div>
                                                        )
                                                }
                                            </td>
                                        </tr>
                                    ))
                                }
                                {/* <tr>
                                <td>
                                    <div className="flex items-center space-x-3">
                                    <div className="avatar">
                                    <div className="mask mask-squircle w-12 h-12">
                                    <img src="https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b" alt="Avatar Tailwind CSS Component" />
                                    </div>
                                    </div>
                                    <div>
                                    <Link to="">
                                    <div className="font-bold">Le Nhat Duy</div>
                                    </Link>
                                    </div>
                                    </div>
                                    </td>
                                    <td>69696969</td>
                                    <td>12/03/2023</td>
                                    <td>12/05/2023</td>
                                    <td>Done</td>
                                </tr>*/}
                            </tbody>
                            {/* foot */}
                            {/* <tfoot>
                                <tr>
                                    <th>Name User</th>
                                    <th>ID Book</th>
                                    <th>From</th>
                                    <th>To</th>
                                    <th>Status</th>
                                </tr>
                            </tfoot> */}

                        </table>
                    </div>
                    <div className="modal-action">
                        <label htmlFor="history_book" className="btn">Cancel</label>
                    </div>
                </div>
            </div>
        </>
    )
}

export default HistoryBook