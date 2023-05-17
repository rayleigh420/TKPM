import { useQuery } from "@tanstack/react-query"
import { Link, useParams } from "react-router-dom"
import { historyVersion } from "../../api/manageApi"
import { data } from "autoprefixer"

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]


// status: 
// booked: yellow
// borrowed: red
// returned: green

const HistoryBook = () => {
    const { id } = useParams()

    const { data: books, isLoading, isError } = useQuery({
        queryKey: ['admin', 'history', 'version', id],
        queryFn: () => historyVersion(id)
    })

    console.log(books)
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
                                    books && books?.map(item => (
                                        <tr>
                                            <td>
                                                <div className="flex items-center space-x-3">
                                                    <div className="avatar">
                                                        <div className="mask mask-squircle w-12 h-12">
                                                            <img src={item?.user.avatar} alt="Avatar Tailwind CSS Component" />
                                                        </div>
                                                    </div>
                                                    <div>
                                                        {/* <Link to="/"> */}
                                                        <div className="font-bold">{item?.user.name}</div>
                                                        {/* </Link> */}
                                                    </div>
                                                </div>
                                            </td>
                                            <td>{item?.book_detail_id}</td>
                                            <td>{item?.date_borrowed}</td>
                                            <td>{item?.date_return || item?.reserve_date}</td>
                                            <td>
                                                {

                                                    (item?.status == 'returned') ?
                                                        (

                                                            <div className="badge bg-green-500 text-[#ffffff] font-semibold">Returned</div>
                                                        )
                                                        :
                                                        (
                                                            (item?.status == 'booked') ?
                                                                <div className="badge bg-yellow-500 text-[#ffffff] font-semibold">Booked</div>
                                                                :
                                                                <div className="badge bg-red-500 text-[#ffffff] font-semibold">Borowed</div>

                                                        )
                                                }
                                            </td>
                                        </tr>
                                    ))
                                }
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