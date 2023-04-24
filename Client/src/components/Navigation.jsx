import { useState } from "react"
import Form from "./Form"
import { Link } from "react-router-dom"

const Navigation = ({ setLog }) => {

    const [search, setSearch] = useState('')

    const handleChangeSearch = (e) => {
        setSearch(e.target.value)
    }

    const handleSubmit = (e) => {
        alert(e.target.value)
    }

    const handleOpen = () => {
        setLog(true)
    }

    return (
        <>
            <div className="navbar bg-[#131415] px-[120px]">
                <div className="flex-1">
                    <Link to="/home">
                        <p className="btn normal-case text-[#ffffff] text-[18px] font-semibold leading-[21.6px]">Library</p>
                    </Link>
                </div>
                <div className="flex-none gap-2">
                    <div className="form-control">
                        <input type="text" placeholder="Search" className="input input-bordered rounded-full bg-[#262627] w-[150px] focus:w-[220px] focus:ease-out" onChange={handleChangeSearch} value={search} onSubmit={handleSubmit} />
                    </div>
                    <div className="dropdown dropdown-end">
                        <label tabIndex={0} className="btn btn-ghost btn-circle avatar" htmlFor="my-modal-4" onClick={handleOpen}>
                            <div className="w-10 rounded-full">
                                <img src="https://assets.stickpng.com/images/585e4bf3cb11b227491c339a.png" />
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
        </>
    )
}

export default Navigation