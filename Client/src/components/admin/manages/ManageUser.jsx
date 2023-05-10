import { Link } from "react-router-dom"
import { data } from "../../../pages/home"
import { useContext, useState } from "react"
import useDebounce from "../../../hooks/useDebounce"
import { useQuery } from "@tanstack/react-query"
import { getAllUser } from "../../../api/userApi"
import AuthContext from "../../../context/AuthProvider"

const ManageUser = () => {
    const [search, setSearch] = useState('')
    const debounced = useDebounce(search, 1000)

    const { auth, setAuth } = useContext(AuthContext)

    const { data: users, isLoading, isError } = useQuery({
        queryKey: ['admin', 'users'],
        queryFn: () => getAllUser(auth.token)
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
                <h1 className="mt-[50px] text-[32px] text-[#ffffff] leading-[32px] font-semibold">Manage User</h1>
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
                                <th>Name User</th>
                                <th>Role</th>
                                <th>Action</th>
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
                                                    {/* <Link to="/book"> */}
                                                    <div className="font-bold">{item.name}</div>
                                                    {/* </Link> */}
                                                </div>
                                            </div>
                                        </td>
                                        <td>User</td>
                                        <td>
                                            <div className="flex flex-row">
                                                {/* <span className="cursor-pointer">
                                                    <EditBook />
                                                    <label htmlFor="modal_editBook" className="cursor-pointer">
                                                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10" />
                                                        </svg>
                                                    </label>
                                                </span> */}
                                                <span className="cursor-pointer">
                                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                        <path strokeLinecap="round" strokeLinejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
                                                    </svg>
                                                </span>
                                            </div>
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
                                        <td>User</td>
                                        <td>
                                            <div className="flex flex-row">
                                                {/* <span className="cursor-pointer">
                                                    <EditBook />
                                                    <label htmlFor="modal_editBook" className="cursor-pointer">
                                                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10" />
                                                        </svg>
                                                    </label>
                                                </span> */}
                                                <span className="cursor-pointer">
                                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                        <path strokeLinecap="round" strokeLinejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
                                                    </svg>
                                                </span>
                                            </div>
                                        </td>
                                    </tr>
                                ))
                            }
                        </tbody>
                        {/* foot */}
                        <tfoot className="sticky bottom-0">
                            <tr>
                                <th>Name User</th>
                                <th>Role</th>
                                <th>Action</th>
                            </tr>
                        </tfoot>
                    </table>
                </div>
            </div >
        </>
    )
}

export default ManageUser