import { useState } from "react"
import BookList from "../components/BookList"
import { data } from "./home"

const Book = () => {
    const [tab, setTab] = useState(false)

    const handleTabAbout = () => {
        setTab(false)
    }

    const handleTabLocation = () => {
        setTab(true)
    }

    return (
        <>
            <div className="flex flex-col items-center">
                <div className=" w-[62%] py-[60px] flex flex-row">
                    <div className="w-[240px] h-[350px] bg-contain hover:border-[1.75px] bg-[url('https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b')] rounded-[7px]">
                    </div>
                    <div className="ml-[30px]">
                        <div className="badge bg-[#eeeeee] rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Ongoing</div>
                        <h1 className="mt-[10px] text-[#ffffff] text-[32px] font-bold leading-[36.8px] tracking-[-0.64]">I Became the 1st Floor Boss of the Tower</h1>
                        <p className="mt-[20px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Author: Si Reubereu</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Type: Light Novel</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Pages: 100</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Year: 2023</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Producer: Wuxiaworld</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Publishing Location: Korean</p>
                        <button className="btn w-[150px] mt-[55px] bg-gradient-to-r from-indigo-700 to-blue-700 text-[#ffffff] leading-[24px] hover:from-indigo-600 hover:to-blue-600">Rent Book</button>
                    </div>
                </div>
                <div className="tabs w-full tab-bordered flex flex-col items-center">
                    <div className="w-[62%]">
                        <a className={"pb-[40px] text-[#ffffff] text-[20px] font-semibold leading-[24px] tracking-[-0.4px] tab tab-bordered" + (tab ? "tab-active" : "")} onClick={handleTabAbout}>About</a>
                        <a className={"pb-[40px] text-[#ffffff] text-[20px] font-semibold leading-[24px] tracking-[-0.4px] tab tab-bordered" + (!tab ? "tab-active" : "")} onClick={handleTabLocation}>Location</a>
                    </div>
                </div>
                {
                    !tab ? (
                        <div className="w-[62%] mt-[30px]">
                            <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px] mb-[20px]">Information</h4>
                            <div className="overflow-x-auto">
                                <table className="table w-full border-nones">
                                    <thead>
                                        <tr>
                                            <th>Pages</th>
                                            <th>Licensed from</th>
                                            <th>Number of renter</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>83 Chapters</td>
                                            <td>Si Reubereu / TITAN</td>
                                            <td>100</td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                            <div className="divider bordered border-[#ffffff]"></div>
                            <div className="text-[#e0e0e0] leading-[24px] tracking-[-0.32px]">
                                <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px] mb-[25px]">Details</h4>
                                Status in Korean: Completed @ 254 chapters <br />
                                <br />
                                Translator: Yeniverse <br />
                                <br />
                                Editor: Dot, SoSam <br />
                                <br />
                                Release rate: 5 chapters a week with an occasional bonus 6th chapter
                            </div>
                            <div className="divider bordered border-[#ffffff]"></div>
                        </div>) :
                        (
                            <div className="w-[62%] mt-[30px] mb-[150px]">
                                <div className="overflow-x-auto">
                                    <table className="table w-full border-nones">
                                        <thead className="text-[#ffffff] font-bold text-[15px]">
                                            <tr>
                                                <th>ID Book</th>
                                                <th>Location</th>
                                                <th>Status</th>
                                            </tr>
                                        </thead>
                                        <tbody className="font-semibold text-[13px]">
                                            <tr>
                                                <td>ABCD6969</td>
                                                <td>123456</td>
                                                <td ><div className="badge badge-info gap-2 w-[50px]">O</div></td>
                                            </tr>
                                            <tr>
                                                <td>ABCD6970</td>
                                                <td>123457</td>
                                                <td><div className="badge badge-warning gap-2 w-[50px]">~</div></td>
                                            </tr>
                                            <tr>
                                                <td>ABCD6971</td>
                                                <td>123458</td>
                                                <td><div className="badge badge-warning gap-2 w-[50px]">X</div></td>
                                            </tr>
                                            <tr>
                                                <td>ABCD6972</td>
                                                <td>123459</td>
                                                <td><div className="badge badge-error gap-2 w-[50px]">X</div></td>
                                            </tr>
                                        </tbody>
                                    </table>
                                    <div className="divider bordered border-[#ffffff]"></div>
                                </div>
                            </div>
                        )
                }
                <div className="w-[62%]">
                    <BookList nameList="Related Book" data={data} />
                </div>
            </div >
            <div className="divider bordered border-[#ffffff]"></div>
        </>
    )
}

export default Book 