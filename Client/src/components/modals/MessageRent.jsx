const MessageRent = ({ rent }) => {
    console.log(rent)
    return (
        <>
            <input type="checkbox" id="messagerent" className="modal-toggle" />
            <label htmlFor="messagerent" className="modal cursor-pointer">
                <label className="modal-box relative" htmlFor="">
                    {!rent ?
                        (
                            <p className="text-center">Please wait!</p>
                        )
                        :
                        (
                            rent == "already borrowing another book"
                                ?
                                (
                                    <p className="text-center capitalize">{rent}</p>
                                )
                                :
                                (
                                    <>
                                        <h3 className="text-lg font-bold text-center">Congratulations you booked this book!</h3>
                                        <p className="py-4 mt-[20px]">Please show this code for librarian to borrowed book.</p>
                                        <p className="text-center text-[20px]">
                                            <kbd className="kbd text-center">{rent}</kbd>
                                        </p>
                                        <p className="text-red-500 text-[14px] mt-[30px]">Remember, This code will be expired in 24h</p>
                                    </>
                                )
                        )
                    }

                </label>
            </label>
        </>
    )
}

export default MessageRent