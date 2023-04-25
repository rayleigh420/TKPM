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
        // setLog(true)
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
                        <label tabIndex={0} className="btn btn-ghost btn-circle avatar" htmlFor="" onClick={handleOpen}>
                            <div className="w-10 rounded-full">
                                <img src="https://assets.stickpng.com/images/585e4bf3cb11b227491c339a.png" />
                            </div>
                        </label>
                        <ul tabIndex={0} className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52 font-semibold">
                            <li><a className="text-[15px] font-bold py-[16px] text-[#ffffff]">
                                Le Nhat Duy
                            </a></li>
                            <li>
                                <Link to="/profile">
                                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                                        <path fillRule="evenodd" d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.25 8.25 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z" clipRule="evenodd" />
                                    </svg>
                                    Profile
                                </Link>
                            </li>
                            <li>
                                <a>
                                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                                        <path fillRule="evenodd" d="M12 6.75a5.25 5.25 0 016.775-5.025.75.75 0 01.313 1.248l-3.32 3.319c.063.475.276.934.641 1.299.365.365.824.578 1.3.64l3.318-3.319a.75.75 0 011.248.313 5.25 5.25 0 01-5.472 6.756c-1.018-.086-1.87.1-2.309.634L7.344 21.3A3.298 3.298 0 112.7 16.657l8.684-7.151c.533-.44.72-1.291.634-2.309A5.342 5.342 0 0112 6.75zM4.117 19.125a.75.75 0 01.75-.75h.008a.75.75 0 01.75.75v.008a.75.75 0 01-.75.75h-.008a.75.75 0 01-.75-.75v-.008z" clipRule="evenodd" />
                                    </svg>
                                    Settings
                                </a>
                            </li>
                            <li>
                                <Link to='/home'>
                                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                                        <path fillRule="evenodd" d="M7.5 3.75A1.5 1.5 0 006 5.25v13.5a1.5 1.5 0 001.5 1.5h6a1.5 1.5 0 001.5-1.5V15a.75.75 0 011.5 0v3.75a3 3 0 01-3 3h-6a3 3 0 01-3-3V5.25a3 3 0 013-3h6a3 3 0 013 3V9A.75.75 0 0115 9V5.25a1.5 1.5 0 00-1.5-1.5h-6zm5.03 4.72a.75.75 0 010 1.06l-1.72 1.72h10.94a.75.75 0 010 1.5H10.81l1.72 1.72a.75.75 0 11-1.06 1.06l-3-3a.75.75 0 010-1.06l3-3a.75.75 0 011.06 0z" clipRule="evenodd" />
                                    </svg>

                                    Logout
                                </Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </div >
        </>
    )
}

export default Navigation