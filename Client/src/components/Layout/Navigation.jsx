import { useContext, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { avaURL } from "../../pages/UserLayout"
import AuthContext from "../../context/AuthProvider"
import { useQuery } from "@tanstack/react-query"
import { getListType } from "../../api/typeApi"
import useLogout from "../../hooks/useLogout"

const Navigation = ({ setLog }) => {

    const [search, setSearch] = useState('')
    const { auth, setAuth } = useContext(AuthContext)
    const navigate = useNavigate()

    const setLogout = useLogout(false)
    // console.log(auth)

    const { data: types, isLoading, iseError } = useQuery({
        queryKey: ['types'],
        queryFn: () => getListType(),
    })

    const handleChangeSearch = (e) => {
        setSearch(e.target.value)
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        if (search == '') {
            return;
        }
        const keyword = encodeURIComponent(search)
        navigate(`/search?keyword=${keyword}`);
        setSearch('')
        console.log(search)
    }

    const handleOpen = () => {
        if (!auth?.name) {
            setLog(true)
        }
    }

    return (
        <>
            <div className="navbar bg-[#131415] px-[120px] fixed top-0 z-50">
                <div className="flex-1 flex flex-row gap-[25px]">
                    <div>
                    </div>
                    <Link to="/" className="pr-[30px]">
                        <p className="normal-case text-[#ffffff] text-[18px] font-semibold leading-[21.6px]">Rayleigh-BaguetteEnjoyer</p>
                    </Link>
                    <Link to="newest" className="pr-[12px]">
                        <p className="normal-case text-[#ffffff] text-[16px] font-bold leading-[19.2px] tracking-[-0.32px]">Newest</p>
                    </Link>
                    <Link to="popular" className="pr-[12px]">
                        <p className="normal-case text-[#ffffff] text-[16px] font-bold leading-[19.2px] tracking-[-0.32px]">Popular</p>
                    </Link>
                    <div className="dropdown dropdown-end">
                        <label tabIndex={0} className="cursor-pointer normal-case text-[#ffffff] text-[16px] font-bold leading-[19.2px] tracking-[-0.32px]" htmlFor="">
                            <div className="">Generics</div>
                        </label>
                        <ul tabIndex={0} className="bg-[#202020] mt-3 p-2 shadow menu menu-compact dropdown-content rounded-box w-52 font-semibold">
                            {
                                types && types.map(item =>
                                (
                                    <li key={item.typeid}>
                                        <Link to={item.typename} className="capitalize">
                                            {item.typename}
                                        </Link>
                                    </li>
                                ))
                            }
                            {/* <li>
                                <Link to="novel">
                                    Novel
                                </Link>
                            </li>
                            <li>
                                <Link to="education">
                                    Education
                                </Link>
                            </li>
                            <li>
                                <Link to='IT'>
                                    Information Technology
                                </Link>
                            </li> */}
                        </ul>
                    </div>
                    {/* <div>
                    </div> */}
                </div>
                <div className="flex-none gap-2">
                    <form onSubmit={handleSubmit}>
                        <div className="form-control">
                            <input type="text" placeholder="Search" className="input input-bordered rounded-full bg-[#262627] w-[150px] focus:w-[220px] focus:ease-out duration-300" onChange={handleChangeSearch} value={search} />
                        </div>
                    </form>
                    <div className="dropdown dropdown-end">
                        <label tabIndex={auth?.name && 0} className="btn btn-ghost btn-circle avatar" onClick={handleOpen}>
                            <div className="w-10 rounded-full">
                                {
                                    auth?.ava ?
                                        (
                                            <img src={auth.ava} />
                                        )
                                        :
                                        (
                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-full h-full overflow-hidden">
                                                <path fillRule="evenodd" d="M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" clipRule="evenodd" />
                                            </svg>

                                        )
                                }
                                {/* <img src={auth?.ava ? auth.ava : "https://assets.stickpng.com/images/585e4bf3cb11b227491c339a.png"} /> */}
                                {/* <img src="https://i1.sndcdn.com/avatars-0QCRofC3yRV0mkpa-6XQLMA-t500x500.jpg" /> */}
                            </div>
                        </label>
                        {auth?.name &&
                            <ul tabIndex={0} className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52 font-semibold">
                                <li>
                                    <Link to="/user/profile" className="text-[15px] font-bold py-[16px] text-[#ffffff]">

                                        <div className="avatar">
                                            <div className="w-7 rounded-full ring ring-[#413ACB]">
                                                {
                                                    auth?.ava ?
                                                        (
                                                            <img src={auth.ava} />
                                                        )
                                                        :
                                                        (
                                                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-full h-full overflow-hidden">
                                                                <path fillRule="evenodd" d="M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" clipRule="evenodd" />
                                                            </svg>

                                                        )
                                                }
                                                {/* <img src={auth.ava ? auth.ava : "https://assets.stickpng.com/images/585e4bf3cb11b227491c339a.png"} /> */}

                                                {/* <img src="https://i1.sndcdn.com/avatars-0QCRofC3yRV0mkpa-6XQLMA-t500x500.jpg" /> */}
                                            </div>
                                        </div>
                                        {auth.name}
                                    </Link>
                                </li>
                                <li>
                                    <Link to="user/profile">
                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                                            <path fillRule="evenodd" d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.25 8.25 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z" clipRule="evenodd" />
                                        </svg>
                                        Profile
                                    </Link>
                                </li>
                                <li>
                                    {
                                        auth?.role == 'admin' ?
                                            (
                                                <Link to="user/manage/viewer">
                                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                        <path strokeLinecap="round" strokeLinejoin="round" d="M10.125 2.25h-4.5c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125v-9M10.125 2.25h.375a9 9 0 019 9v.375M10.125 2.25A3.375 3.375 0 0113.5 5.625v1.5c0 .621.504 1.125 1.125 1.125h1.5a3.375 3.375 0 013.375 3.375M9 15l2.25 2.25L15 12" />
                                                    </svg>

                                                    Manage
                                                </Link>
                                            )
                                            :
                                            (
                                                <Link to="user/listrented">
                                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                        <path strokeLinecap="round" strokeLinejoin="round" d="M9 12h3.75M9 15h3.75M9 18h3.75m3 .75H18a2.25 2.25 0 002.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 00-1.123-.08m-5.801 0c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75 2.25 2.25 0 00-.1-.664m-5.8 0A2.251 2.251 0 0113.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m0 0H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V9.375c0-.621-.504-1.125-1.125-1.125H8.25zM6.75 12h.008v.008H6.75V12zm0 3h.008v.008H6.75V15zm0 3h.008v.008H6.75V18z" />
                                                    </svg>

                                                    List rented book
                                                </Link>
                                            )
                                    }
                                </li>
                                <li onClick={() => setLogout(true)}>
                                    <Link to='/'>
                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                                            <path fillRule="evenodd" d="M7.5 3.75A1.5 1.5 0 006 5.25v13.5a1.5 1.5 0 001.5 1.5h6a1.5 1.5 0 001.5-1.5V15a.75.75 0 011.5 0v3.75a3 3 0 01-3 3h-6a3 3 0 01-3-3V5.25a3 3 0 013-3h6a3 3 0 013 3V9A.75.75 0 0115 9V5.25a1.5 1.5 0 00-1.5-1.5h-6zm5.03 4.72a.75.75 0 010 1.06l-1.72 1.72h10.94a.75.75 0 010 1.5H10.81l1.72 1.72a.75.75 0 11-1.06 1.06l-3-3a.75.75 0 010-1.06l3-3a.75.75 0 011.06 0z" clipRule="evenodd" />
                                        </svg>
                                        Logout
                                    </Link>
                                </li>
                            </ul>
                        }
                    </div>
                </div>
            </div >
        </>
    )
}

export default Navigation