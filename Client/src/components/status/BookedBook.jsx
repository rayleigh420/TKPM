import { Link } from "react-router-dom"
import { data } from "../../pages/home"
import AddBook from "../modals/AddBook"
import EditBook from "../modals/EditBook"
import { useState } from "react"

const BookedBook = () => {
    const [search, setSearch] = useState('')

    const handleChangeSeach = (e) => {
        setSearch(e.target.value)
    }

    const handleSubmitSearch = (e) => {
        e.preventDefault()
        console.log(search, typeSearch)
    }

    const handleChangeBorrowed = (index) => {
        console.log("Change to borrowed: ", index)
    }

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
                                <th>From</th>
                                <th>To</th>
                            </tr>
                        </thead>
                        <tbody>
                            {/* row 1 */}
                            {
                                data.map((item, index) => (
                                    <tr key={index}>
                                        <td>
                                            <label>
                                                <input type="checkbox" className="checkbox" onChange={() => handleChangeBorrowed(index)} />
                                            </label>
                                        </td>
                                        <td>
                                            <div className="flex items-center space-x-3">
                                                <div className="avatar">
                                                    <div className="mask mask-squircle w-12 h-12">
                                                        <img src={item.src} alt="Avatar Tailwind CSS Component" />
                                                    </div>
                                                </div>
                                                <div>
                                                    {/* <Link to="/book"> */}
                                                    <div className="font-bold">{item.name}</div>
                                                    {/* </Link> */}
                                                </div>
                                            </div>
                                        </td>
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
                                                        <div className="text-sm opacity-50">69696969</div>
                                                    </Link>
                                                </div>
                                            </div>
                                        </td>
                                        <td>27/03/2023</td>
                                        <td>27/04/2023</td>
                                    </tr>
                                ))
                            }
                            {
                                data.map((item, index) => (
                                    <tr key={index}>
                                        <td>
                                            <label>
                                                <input type="checkbox" className="checkbox" onChange={() => handleChangeBorrowed(index)} />
                                            </label>
                                        </td>
                                        <td>
                                            <div className="flex items-center space-x-3">
                                                <div className="avatar">
                                                    <div className="mask mask-squircle w-12 h-12">
                                                        <img src={item.src} alt="Avatar Tailwind CSS Component" />
                                                    </div>
                                                </div>
                                                <div>
                                                    {/* <Link to="/book"> */}
                                                    <div className="font-bold">{item.name}</div>
                                                    {/* </Link> */}
                                                </div>
                                            </div>
                                        </td>
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
                                                        <div className="text-sm opacity-50">69696969</div>
                                                    </Link>
                                                </div>
                                            </div>
                                        </td>
                                        <td>27/03/2023</td>
                                        <td>27/04/2023</td>
                                    </tr>
                                ))
                            }
                        </tbody>
                        {/* foot */}
                        <tfoot className="sticky bottom-0">
                            <tr>
                                <th>Borrowed</th>
                                <th>User</th>
                                <th>Book</th>
                                <th>From</th>
                                <th>To</th>
                            </tr>
                        </tfoot>
                    </table>
                </div>
            </div >
        </>
    )
}

export default BookedBook
