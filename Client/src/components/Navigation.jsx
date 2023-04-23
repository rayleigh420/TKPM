import { useState } from "react"

const Navigation = () => {

    const [search, setSearch] = useState('')

    const handleChangeSearch = (e) => {
        setSearch(e.target.value)
    }

    const handleSubmit = (e) => {
        alert(e.target.value)

    }

    return (
        <div className="navbar bg-[#131415] px-[120px]">
            <div className="flex-1">
                <a className="btn btn-ghost normal-case text-[#ffffff] text-[18px] font-semibold leading-[21.6px]">Library</a>
            </div>
            <div className="flex-none gap-2">
                <div className="form-control">
                    <input type="text" placeholder="Search" className="input input-bordered rounded-full bg-[#262627] w-[150px] focus:w-[220px] focus:ease-out" onChange={handleChangeSearch} value={search} onSubmit={handleSubmit} />
                </div>
                <div className="dropdown dropdown-end">
                    <label tabIndex={0} className="btn btn-ghost btn-circle avatar" htmlFor="my-modal-4">
                        <div className="w-10 rounded-full">
                            <img src="/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                        </div>
                    </label>
                    {/* <ul tabIndex={0} className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                        <li>
                            <a className="justify-between">
                                Profile
                                <span className="badge">New</span>
                            </a>
                        </li>
                        <li><a>Settings</a></li>
                        <li><a>Logout</a></li>
                    </ul> */}
                </div>
            </div>
        </div>
    )
}

export default Navigation