import { useMutation, useQuery } from "@tanstack/react-query"
import AddVersionBook from "../modals/AddVersionBook"
import HistoryBook from "./HistoryBook"
import { getVersionBook } from "../../api/bookApi"
import { addVersionBook } from "../../api/manageApi"
import { toast } from "react-toastify"

const LocationVersion = ({ id }) => {

    const { data: version, isLoading, isError } = useQuery({
        queryKey: ['version', id],
        queryFn: () => getVersionBook(id)
    })

    console.log(version)

    return (
        <div className="w-[62%] mt-[30px] mb-[50px]">
            <div className="">
                <div className="flex flex-row justify-between mb-[20px]">
                    <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px]">Status of all versions</h4>
                    <div>
                        <HistoryBook id={id} />
                        {/* <button className="btn "> */}
                        <label htmlFor="history_book" className="cursor-pointer btn bg-[#eeeeee] text-[#000000] mr-[10px] hover:text-[#ffffff]">
                            Watch History Rent
                        </label>
                        {/* </button> */}
                        <AddVersionBook />
                        {/* <button className=""> */}
                        <label htmlFor="modal_addVersion" className="cursor-pointer btn w-[140px] bg-gradient-to-r from-indigo-700 to-blue-700 text-[#ffffff] leading-[24px] hover:from-indigo-600 hover:to-blue-600">
                            Add version
                        </label>
                        {/* </button> */}
                    </div>

                </div>
                <table className="table w-full border-nones">
                    <thead className="text-[#ffffff] font-bold text-[15px]">
                        <tr>
                            <th>ID Book</th>
                            <th>Location</th>
                            <th>Status</th>
                        </tr>
                    </thead>
                    <tbody className="font-semibold text-[13px]">
                        {
                            version && version.map(item => (
                                <tr key={item._id}>
                                    <td>{item.book_detail_id}</td>
                                    <td>{item.location}</td>
                                    <td>
                                        <div className={"badge gap-2 w-[80px] text-[#ffffff] capitalize " + (item.status == 'ready' ? "bg-green-500" : (item.status == 'booked' ? 'bg-yellow-500' : 'bg-red-500'))}>
                                            {item.status}
                                        </div>
                                    </td>
                                </tr>
                            ))
                        }
                        {/* <tr>
                            <td>ABCD6969</td>
                            <td>123456</td>
                            <td ><div className="badge bg-green-500 gap-2 w-[80px] text-[#ffffff]">Available</div></td>
                        </tr>
                        <tr>
                            <td>ABCD6970</td>
                            <td>123457</td>
                            <td><div className="badge bg-yellow-500 gap-2 w-[80px] text-[#ffffff]">Booked</div></td>
                        </tr>
                        <tr>
                            <td>ABCD6971</td>
                            <td>123458</td>
                            <td><div className="badge bg-red-500 gap-2 w-[80px] text-[#ffffff]">Borrowed</div></td>
                        </tr>
                        <tr>
                            <td>ABCD6972</td>
                            <td>123459</td>
                            <td><div className="badge bg-red-500 gap-2 w-[80px] text-[#ffffff]">Borrowed</div></td>
                        </tr> */}
                    </tbody>
                </table>
                <div className="divider bordered border-[#ffffff]"></div>
            </div>
        </div>
    )
}

export default LocationVersion