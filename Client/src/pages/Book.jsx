const Book = () => {
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
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Producer: Wuxiaworld</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Publishing Location: Korean</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Year: 2023</p>
                        <p className="mt-[10px] text-[#bdbdbd] text-[15.2px] font-medium leading-[18.24px] tracking-[-0.304px]">Type: Light Novel</p>
                        <button className="btn w-[150px] mt-[82px] bg-gradient-to-r from-indigo-700 to-blue-700 text-[#ffffff] leading-[24px] hover:from-indigo-500 hover:to-blue-500">Rent Book</button>
                    </div>
                </div>
            </div >
            {/* <img src="" /> */}
        </>
    )
}

export default Book 