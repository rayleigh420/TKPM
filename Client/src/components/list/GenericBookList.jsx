const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]

const GenericBookList = () => {
    return (
        <div className="flex flex-col items-start mt-[50px]">
            <p className='text-[28px] font-bold leading-[32.2px] tracking-[-0.56px] text-[#ffffff]'>Generics</p>
            <div class="w-full mt-[20px] grid grid-cols-2 gap-y-6 gap-x-3">
                {
                    arr.map(item => (
                        <div className="flex flex-row gap-[12px]">
                            <div className="cursor-pointer min-w-[125px] h-[180px] bg-cover hover:border-[0.1px] hover:border-[#142B45] bg-[url('https://cdn.wuxiaworld.com/images/covers/bfbt.jpg?ver=fbc27beb0a7017f23af5a9560099d3aeaa2b8d2b')] rounded-[7px]"></div>
                            <div className="">
                                <div className="badge bg-[#eeeeee] rounded-[4px] text-[10.4px] font-bold leading-[12.48px] tracking-[-0.208] text-[#000000]">Ongoing</div>
                                <h1 className="mt-[5px] text-[#ffffff] font-bold leading-[20px]">I Became the 1st Floor Boss of the Tower</h1>
                                <p className="mt-[10px] text-[#bdbdbd] text-[14px] font-medium">Author: Si Reubereu</p>
                                <p className="text-ellipsis overflow-hidden mt-[10px] text-[#bdbdbd] text-[14px] font-medium">
                                    For decades, he had lived as a puppet of the tower. Throughout that time, who knew how many deaths he had encountered? One day, however, his memories returned to him. The First Floor Boss finally awakened
                                </p>
                            </div>
                        </div>
                    ))
                }
            </div>
        </div>
    )
}

export default GenericBookList