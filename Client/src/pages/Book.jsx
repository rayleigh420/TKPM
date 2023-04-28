import { useState } from "react"
import { data } from "./home"
import HistoryBook from "../components/HistoryBook"
import AddVersionBook from "../components/modals/AddVersionBook"
import RelatedBook from "../components/RelatedBook"

//status:

// red: borrowed
// yellow: booked
// green: available

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
                    <div className="w-[240px] h-[350px] bg-cover hover:border-[1.75px] bg-[url('https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b')] rounded-[7px]">
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
                            <div className="text-[#e0e0e0] leading-[24px] tracking-[-0.32px]">
                                <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px] mb-[25px]">Description</h4>
                                For decades, he had lived as a puppet of the tower.
                                Throughout that time, who knew how many deaths he had encountered?
                                One day, however, his memories returned to him.
                                The First Floor Boss finally awakened
                            </div>
                            <div className="divider bordered border-[#ffffff]"></div>
                        </div>) :
                        (
                            <div className="w-[62%] mt-[30px] mb-[150px]">
                                <div className="">
                                    <div className="flex flex-row justify-between mb-[20px]">
                                        <h4 className="text-[#ffffff] text-[20px] font-bold leading-[24px] tracking-[-0.4px]">Status of all versions</h4>
                                        <div>
                                            <HistoryBook />
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
                                            <tr>
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
                                            </tr>
                                        </tbody>
                                    </table>
                                    <div className="divider bordered border-[#ffffff]"></div>
                                </div>
                            </div>
                        )
                }
                <div className="w-[62%]">
                    <RelatedBook />
                </div>
            </div >
            <div className="divider bordered border-[#ffffff]"></div>
        </>
    )
}

export default Book 