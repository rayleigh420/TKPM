package seeds

import (
	"context"
	"time"

	"github.com/baguette/seed-lib/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var now, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
var week1, _ = time.Parse(time.RFC3339, time.Now().Add(-7*24*time.Hour).Format(time.RFC3339))

func SeedBook() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	Book1 := models.BookModel{
		Name:                "Không Diệt Không Sinh Đừng Sợ Hãi",
		Publisher:           "Saigon Books",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/352/200/300.jpg?hmac=JRE6d4eB1tvPUpBESG8XEM2_22EaXNe2luRrVkydr2E",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   12,
		Amount:              5,
		Type_id:             "2",
		Publishing_location: "VietNam",
		Page:                100,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book2 := models.BookModel{
		Name:                "Muôn Kiếp Nhân Sinh - Many Times, Many Lives",
		Publisher:           "NXB Tổng Hợp TPHCM",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/94/200/300.jpg?hmac=CA7x5S3EdSeRqM9TK0RxiKTcx1R96lIncvKTjTP3beI",
		Author:              "Nguyên Phong",
		Borrowed_quantity:   120,
		Amount:              5,
		Type_id:             "2",
		Publishing_location: "VietNam",
		Page:                100,
		License:             "Nguyên Phong",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book3 := models.BookModel{
		Name:                "Muốn An Được An",
		Publisher:           "Thái Hà",
		Yearpublished:       2021,
		Book_img:            "https://fastly.picsum.photos/id/882/200/300.jpg?hmac=j02hkJZfREXyROkSfM6KpIIPHsAaGLmWylVEr5sRzok",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   20,
		Amount:              5,
		Type_id:             "2",
		Publishing_location: "VietNam",
		Page:                100,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book4 := models.BookModel{
		Name:                "Tĩnh Lặng - Sức Mạnh Tĩnh Lặng Trong Thế Giới Huyền Ảo",
		Publisher:           "Thái Hà",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/25/200/300.jpg?hmac=ScdLbPfGd_kI3MUHvJUb12Fsg1meDQEaHY_mM613BVM",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   42,
		Amount:              5,
		Type_id:             "2",
		Publishing_location: "Korea",
		Page:                100,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book5 := models.BookModel{
		Name:                "Chainsaw Man - Tập 11",
		Publisher:           "NXB Trẻ",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/200/200/300.jpg?hmac=XVCLpc2Ddr652IrKMt3L7jISDD4au5O9ZIr3fwBtxo8",
		Author:              "Tatsuki Fujimoto",
		Borrowed_quantity:   512,
		Amount:              5,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                100,
		License:             "Tatsuki Fujimoto",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book6 := models.BookModel{
		Name:                "Khu Rừng Đom Đóm",
		Publisher:           "NXB Trẻ",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/461/200/300.jpg?hmac=dIwmQxeVflRD0QrOZ_p0_q-LpAY7sVhua6FCEIR_xi8",
		Author:              "Yuki Midorikawa",
		Borrowed_quantity:   162,
		Amount:              5,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                300,
		License:             "Yuki Midorikawa",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book7 := models.BookModel{
		Name:                "Cô Gái Nơi Xứ Ngoài",
		Publisher:           "IPM",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/615/200/300.jpg?hmac=ehJCfeXO1-ZbwBXgbYKroA97kTtoPKNoyEbCxnzsYfU",
		Author:              "Nagabe",
		Borrowed_quantity:   127,
		Amount:              5,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                352,
		License:             "Nagabe",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book8 := models.BookModel{
		Name:                "Cô Gái Nơi Xứ Ngoài - Tập 2",
		Publisher:           "IPM",
		Yearpublished:       2001,
		Book_img:            "https://fastly.picsum.photos/id/183/200/300.jpg?hmac=Z9yCtuuIPn5CuOhwIntNEQFIRotghuBn06nqOSL828c",
		Author:              "Nagabe",
		Borrowed_quantity:   182,
		Amount:              5,
		Type_id:             "1",
		Publishing_location: "Japan",
		Page:                451,
		License:             "Nagabe",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book9 := models.BookModel{
		Name:                "Blockchain - Bản Chất Của Blockchain",
		Publisher:           "1980 Books",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/435/200/300.jpg?hmac=gSfIK7r5rB4f3aJq0RTylep967d4sBRbIUuAOuq433o",
		Author:              "Mark Gates",
		Borrowed_quantity:   192,
		Amount:              5,
		Type_id:             "3",
		Publishing_location: "USA",
		Page:                256,
		License:             "Mark Gates",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book10 := models.BookModel{
		Name:                "Object Oriented Programming",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/972/200/300.jpg?hmac=UMf5f6BV9GkLiz0Xz9kMwm1riiTtlpIG2jt0WrxZ51Q",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   1012,
		Amount:              5,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                444,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book11 := models.BookModel{
		Name:                "Sống Chết Mỗi Ngày: Hành Trình Đi Xuyên Qua Các Tiến Trình Sinh Tử Của Một Nhà Sư Phật Giáo",
		Publisher:           "Nhà Xuất Bản Hà Nội",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/60/1920/1200.jpg?hmac=fAMNjl4E_sG_WNUjdU39Kald5QAHQMh-_-TsIbbeDNI",
		Author:              "Yongey Mingyur Rinpoche",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4", //TonGiao - Tam Linh
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Yongey Mingyur Rinpoche, sinh năm 1975 tại Nubri, Nepal, là con trai út của bậc thầy, thiền giả danh tiếng Tulku Urgyen Rinpoche. Ngài bắt đầu tu học chính thức ở tuổi mười một và hai năm sau bắt đầu khóa nhập thất ba năm đầu tiên của mình. Ngày nay, giáo lý của ngài tích hợp tính thực tiễn và quy luật triết học của các chương trình đào tạo Tây Tạng với các định hướng khoa học và tâm lý học của phương Tây. Ngoài vai trò là trụ trì của ba tu viện, ngài cũng lãnh đạo Tergar, một cộng đồng thiền quốc tế với một trăm trung tâm trên khắp thế giới. Ngài được biết rộng rãi đến nhờ cách trình bày thực hành thiền định rất rõ ràng và dễ tiếp cận. `,
		Description:         `Cuốn sách “Sống chết mỗi ngày: Hành trình đi xuyên qua các tiến trình Sinh Tử của một nhà sư Phật giáo” đã được Pema Chodron - tác giả của cuốn “Khi mọi thứ sụp đổ” đã có lời khen tặng dành cho cuốn sách: “Thông qua cuốn sách này, chúng ta được bước vào thế giới nội tâm của một nhà sư Phật giáo trẻ tuổi và kiệt xuất.`,
	}
	Book12 := models.BookModel{
		Name:                "Cây Cam Ngọt Của Tôi",
		Publisher:           "Nhà Xuất Bản Hội Nhà Văn",
		Yearpublished:       2010,
		Book_img:            "https://fastly.picsum.photos/id/61/3264/2448.jpg?hmac=Xi3Kq99U5tueFi0aajgUP0yWAL4xwCHg5ZXGZRLTqyI",
		Author:              "José Mauro de Vasconcelos",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "5",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Nhã Nam",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book13 := models.BookModel{
		Name:                "Gửi Tới Người Mùa Đông Này Sẽ Không Còn Nữa",
		Publisher:           "Nhà Xuất Bản Dân Trí",
		Yearpublished:       2022,
		Book_img:            "https://fastly.picsum.photos/id/62/2000/1333.jpg?hmac=PbFIn8k0AndjiUwpOJcfHz2h-wPCQi_vJRTJZPdr6kQ",
		Author:              "Inujin",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "6",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "ShineBooks",
		Details:             `Một ngày đông nọ, Natsumi Ikuta, cô gái 24 tuổi đang bế tắc trong công việc và cuộc sống đã bị mắc kẹt trong một đám cháy khi tăng ca tại công ty. Ngay lúc cô sẵn sàng đón nhận cái chết, một chàng trai tên Atsuki đột nhiên xuất hiện và cứu sống cô. Song cậu nói rằng cô sẽ đối mặt với cái chết vào thời điểm này của năm trong vòng sáu năm tới, và nhằm ngăn chặn điều đó, mỗi năm Atsuki sẽ xuất hiện để giúp đỡ cô. Liệu Natsumi có thể vượt qua những thử thách của tử thần để tiếp tục sống sót không? Và thân phận thực sự của Atsuki là gì, tại sao cậu lại giúp đỡ cô?`,
		Description:         `Một ngày đông nọ, Natsumi Ikuta, cô gái 24 tuổi đang bế tắc trong công việc và cuộc sống đã bị mắc kẹt trong một đám cháy khi tăng ca tại công ty. Ngay lúc cô sẵn sàng đón nhận cái chết, một chàng trai tên Atsuki đột nhiên xuất hiện và cứu sống cô. Song cậu nói rằng cô sẽ đối mặt với cái chết vào thời điểm này của năm trong vòng sáu năm tới, và nhằm ngăn chặn điều đó, mỗi năm Atsuki sẽ xuất hiện để giúp đỡ cô. Liệu Natsumi có thể vượt qua những thử thách của tử thần để tiếp tục sống sót không? Và thân phận thực sự của Atsuki là gì, tại sao cậu lại giúp đỡ cô?`,
	}
	Book14 := models.BookModel{
		Name:                "Nơi Khu Rừng Chạm Tới Những Vì Sao",
		Publisher:           "Nhà Xuất Bản Thanh Niên",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/63/5000/2813.jpg?hmac=HvaeSK6WT-G9bYF_CyB2m1ARQirL8UMnygdU9W6PDvM",
		Author:              "Glendy Vanderah",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "6",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "1980 Books",
		Details:             `Đứa trẻ bí ẩn xuất hiện với đôi chân trần, trên người đầy vết bầm tím. \n Cô bé tự xưng là Ursa, được cử đến từ những vì tinh tú và sẽ chỉ rời đi sau khi chứng kiến đủ năm điều kỳ diệu nơi Địa Cầu. Lo lắng cho đứa trẻ, Jo - một nhà sinh vật học thực địa sống giữa vùng nông thôn Illinois- đành để cô bé ở lại, chỉ tới khi cô biết được sự thật về quá khứ của đứa bé đó.`,
		Description:         `Những vấn đề rắc rối bắt đầu xuất hiện khi Jo nhờ đến sự giúp đỡ của Gabriel Nash, một anh chàng hàng xóm ẩn dật, để tìm hiểu về câu chuyện bí ẩn của đứa trẻ kia. Nhưng ở bên cô bé càng nhiều, họ càng nảy sinh những thắc mắc. Vì sao một đứa trẻ không chỉ đọc hết, mà còn hiểu rõ những tác phẩm của Shakespeare? Vì sao những điều tốt đẹp luôn xuất hiện cùng cô bé? Và tại sao Jo và Gabe ngừng tìm kiếm thông tin cô bé trên trang web về trẻ em mất tích?`,
	}
	Book15 := models.BookModel{
		Name:                "Khu Vườn Ngôn Từ",
		Publisher:           "Nhà Xuất Bản Văn Học",
		Yearpublished:       2015,
		Book_img:            "https://fastly.picsum.photos/id/64/4326/2884.jpg?hmac=9_SzX666YRpR_fOyYStXpfSiJ_edO3ghlSRnH2w09Kg",
		Author:              "Shinkai Makoto",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "6",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "IBM",
		Details:             `Khu vườn ngôn từ kể về một tình yêu còn xa xưa hơn cả tình yêu. \n Khái niệm tình yêu trong tiếng Nhật hiện đại là luyến hoặc ái, nhưng vào thời xưa nó được viết là cô bi, nghĩa là nỗi buồn một mình. Shinkai Makoto đã cấu tứ Khu vườn ngôn từ theo ý nghĩa cổ điển này, miêu tả tình yêu theo khái niệm ban sơ của nó, tức là cô bi - nỗi buồn khi một mình thương nhớ một người.`,
		Description:         `Làm nền cho tất cả là mưa rơi không ngừng, là lá mướt mát rung rinh. Nhưng khi mưa tạnh và trời quang trở lại, mọi đường nét của hiện thực trở nên rõ rệt đến khắc nghiệt, thì những êm dịu và lửng lơ kia liệu còn khả năng tồn tại?`,
	}
	Book16 := models.BookModel{
		Name:                "Từ Điển Tiếng \"Em\"",
		Publisher:           "Nhà Xuất Bản Phụ Nữ",
		Yearpublished:       2022,
		Book_img:            "https://fastly.picsum.photos/id/65/4912/3264.jpg?hmac=uq0IxYtPIqRKinGruj45KcPPzxDjQvErcxyS1tn7bG0",
		Author:              "Khotudien",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                280,
		License:             "Skybooks",
		Details:             `Bạn sẽ bất ngờ, khi cầm cuốn “từ điển” xinh xinh này trên tay. \nVà sẽ còn ngạc nhiên hơn nữa, khi bắt đầu đọc từng trang sách… \n Dĩ nhiên là vì “Từ điển tiếng “Em” không phải là một cuốn từ điển thông thường rồi!`,
		Description:         `Và bạn yên tâm, “tập đoàn” Kho Từ Điển điều hành bởi Thịt Kho - Hiệp Thị - 2 chủ xị cho ra đời cuốn sách nhỏ bé xíu xiu nhưng mới mẻ và mặn mà vô cùng này sẽ khiến bạn thoát mác “người tối cổ” cười cả ngày không chán, nhìn bạn bè quanh mình bằng ánh mắt dễ thương, tận hưởng cuộc đời với những định nghĩa hoàn toàn!!!`,
	}
	Book17 := models.BookModel{
		Name:                "Giá Như Có Ai Đó Dạy Tôi Trưởng Thành",
		Publisher:           "Nhà Xuất Bản Văn Học",
		Yearpublished:       2022,
		Book_img:            "https://fastly.picsum.photos/id/66/3264/2448.jpg?hmac=H9yvGug9-Lk5f-1qZqs6dEV-Yd40jFOIC7oudo4eBK4",
		Author:              "Bloom",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "7",
		Publishing_location: "VietNam",
		Page:                216,
		License:             "CÔNG TY TNHH VĂN HÓA VÀ TRUYỀN THÔNG THẢO NGUYÊN",
		Details:             `Cuốn sách này dành cho bạn - những đứa trẻ đang tập lớn mang trong mình trái tim tổn thương và rạn vỡ, những người đang phải đối mặt với những nỗi đau hay những cảm xúc tiêu cực bên trong mình. Cuốn sách xinh xắn của tác giả Bloom sẽ như một lời nhắc nhở, một cái nắm tay giúp bạn hiểu rằng mình không hề đơn độc, rằng bạn hoàn toàn có đủ sức mạnh để đối diện với những khó khăn, để yêu thương, để tự chữa lành cho bản thân và dần dần lớn lên.`,
		Description:         `Bởi trong quá trình trưởng thành, đôi khi chúng ta sẽ tự hỏi, rằng liệu mình đang “sống” hay “tồn tại”, mình là ai trong cuộc đời, mình thực sự cần gì và mong muốn gì. Nhưng có lẽ, không phải ai cũng dễ dàng tìm được cho bản thân câu trả lời, không phải ai ngay từ ban đầu cũng đã đủ bản lĩnh để thích nghi với những chặng đường ngày càng thêm khó. Mà chúng ta - những công trình dang dở của Vũ trụ, bằng cách lắng nghe trái tim và tin tưởng vào bản thân mình, cứ thế can đảm bước đi, lần lượt trải nghiệm và vượt qua những thử thách to nhỏ, cho phép mình được sai, được học, được sửa, được làm lại. Để rồi đến một hôm, tất cả những gì chúng ta tìm kiếm, một cách tự nhiên đều vào đúng chỗ của chúng.`,
	}
	Book18 := models.BookModel{
		Name:                "Đọc Vị Tuổi Dậy Thì Và Hội Chứng Tuổi Teen",
		Publisher:           "Nhà Xuất Bản Lao Động",
		Yearpublished:       2016,
		Book_img:            "https://fastly.picsum.photos/id/67/2848/4288.jpg?hmac=X_Z0Wdd3HiJ8eWT0ohdmpRSIA6e6s265INUMuHA8MqA",
		Author:              "Lee Jin Ah",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "7",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "IBM",
		Details:             `Các phụ huynh có con tuổi dậy thì không cần sự dạy bảo hay hướng dẫn một chiều mà họ mong muốn được đồng cảm và động viên. Bọn trẻ rất vất vả khi lần đầu tiên đối diện với tuổi dậy thì, và đối với các phụ huynh đang ở giai đoạn tiền mãn kinh thì đây cũng lần đầu tiên họ đồng hành cùng bọn trẻ đang tuổi dậy thì nên bản thân họ cũng vất vả không kém.`,
		Description:         `Các phụ huynh có con tuổi dậy thì không cần sự dạy bảo hay hướng dẫn một chiều mà họ mong muốn được đồng cảm và động viên. Bọn trẻ rất vất vả khi lần đầu tiên đối diện với tuổi dậy thì, và đối với các phụ huynh đang ở giai đoạn tiền mãn kinh thì đây cũng lần đầu tiên họ đồng hành cùng bọn trẻ đang tuổi dậy thì nên bản thân họ cũng vất vả không kém.`,
	}
	Book19 := models.BookModel{
		Name:                "Dấu Chân Trên Cát",
		Publisher:           "Nhà Xuất Bản Tổng hợp TP.HCM",
		Yearpublished:       2020,
		Book_img:            "https://fastly.picsum.photos/id/68/4608/3072.jpg?hmac=0KfOS12jehkc6HbfMXWj3GMFve9kVs82dYsN12Npn2Y",
		Author:              "Nguyên Phong",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "TFirst News - Trí Việt",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book20 := models.BookModel{
		Name:                "Fear - Sợ Hãi (Hóa Giải Sợ Hãi Bằng Tình Thương)",
		Publisher:           "Thái Hà",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/70/3011/2000.jpg?hmac=-npCfe1kpGYW7HcBlZvrEZ9Qb_EdiGLbDxE26amgotM",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book21 := models.BookModel{
		Name:                "Nhân Duyên Tiền Kiếp",
		Publisher:           "Dick Sutphen",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/972/200/300.jpg?hmac=UMf5f6BV9GkLiz0Xz9kMwm1riiTtlpIG2jt0WrxZ51Q",
		Author:              "Dick Sutphen",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Dick Sutphen's License",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book22 := models.BookModel{
		Name:                "Sách Lời Tiên Tri Celestine - First News",
		Publisher:           "James Redfield",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/72/3000/2000.jpg?hmac=hPLN3NcJrehzDdebypIHkhh2RLnh89HwJ8QemMsRjzc",
		Author:              "James Redfield",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "James Redfield",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book23 := models.BookModel{
		Name:                "Bên Rặng Tuyết Sơn",
		Publisher:           "Nguyên Phong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/73/5000/3333.jpg?hmac=Tp-n4qzO5flXetGX76h2Ht1P6PHBQlabJuULUtj6ytw",
		Author:              "Nguyên Phong",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Nguyên Phong",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book24 := models.BookModel{
		Name:                "Con Đường Chuyển Hóa - Kinh Bốn Lĩnh Vực Quán Niệm",
		Publisher:           "Thich Nhất Hạnh",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/74/4288/2848.jpg?hmac=q02MzzHG23nkhJYRXR-_RgKTr6fpfwRgcXgE0EKvNB8",
		Author:              "Thich Nhất Hạnh",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Thich Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book25 := models.BookModel{
		Name:                "OSHO Từ Bi - Compassion",
		Publisher:           "OSHO",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/76/4912/3264.jpg?hmac=VkFcSa2Rbv0R0ndYnz_FAmw02ON1pPVjuF_iVKbiiV8",
		Author:              "OSHO",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "OSHO",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book26 := models.BookModel{
		Name:                "Giận (Tái Bản)",
		Publisher:           "Thích Nhất Hạnh",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/75/1999/2998.jpg?hmac=0agRZd8c5CRiFvADOWJqfTv6lqYBty3Kw-9LEtLp_98",
		Author:              "Thích Nhất Hạnh",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Thích Nhất Hạnh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book27 := models.BookModel{
		Name:                "Thực Hành Thiền Định",
		Publisher:           "Matthieu Ricard",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/82/1500/997.jpg?hmac=VcdCqu9YiLpbCtr8YowUCSUD3-245TGekiXmtiMXotw",
		Author:              "Matthieu Ricard",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "4",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Matthieu Ricard",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book28 := models.BookModel{
		Name:                "Excel Ứng Dụng Văn Phòng ĐÀO TẠO TIN HỌC Từ Cơ Bản Đến Nâng Cao",
		Publisher:           "Nguyễn Quang Vinh",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/83/2560/1920.jpg?hmac=LFdAxfpbYKs0hZr0LhHVWyqXarWGg7FtM8pIzJPBc0w",
		Author:              "Nguyễn Quang Vinh",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Nguyễn Quang Vinh",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book29 := models.BookModel{
		Name:                "Giáo Trình Kỹ Thuật Lập Trình C Căn Bản Và Nâng Cao",
		Publisher:           "Huy Hoàng Bookstore",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/84/1280/848.jpg?hmac=YFRYDI4UsfbeTzI8ZakNOR98wVU7a-9a2tGF542539s",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book30 := models.BookModel{
		Name:                "Các hình thức tấn công mạng - Cyberspace",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/87/1280/960.jpg?hmac=tyU21LuCEO1qRepY4GnT9gGkfKbvY__ZrZYg_JxZxI8",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book31 := models.BookModel{
		Name:                "Python lập trình thuật toán",
		Publisher:           "Bùi Việt Hà",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/88/1280/1707.jpg?hmac=NnkwPVDBTVxHkc4rALB_fyu-OHY2usdm7iRk5El7JC4",
		Author:              "Bùi Việt Hà",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book32 := models.BookModel{
		Name:                "Khởi Nghiệp Kinh Doanh Online - Bán Hàng Hiệu Quả Trên Facebook",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/89/4608/2592.jpg?hmac=G9E4z5RMJgMUjgTzeR4CFlORjvogsGtqFQozIRqugBk",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "10",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book33 := models.BookModel{
		Name:                "Tâm Lý Học Về Tiền",
		Publisher:           "Morgan Housel",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/81/5000/3250.jpg?hmac=gFiddUc7I25C06HUIMesyaFCjSOWE3L3uDl0QSyuX4M",
		Author:              "Morgan Housel",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "10",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Morgan Housel",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book34 := models.BookModel{
		Name:                "Bí Quyết Tìm Kiếm Mặt Bằng Kinh Doanh - Đến Sahara Mở Quán Trà Đá - Bộ Sách Khởi Nghiệp Bán Lẻ",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/80/3888/2592.jpg?hmac=zD95NwXZ7mGAMj-z4444Elf43I4HJvd7Afm2tloweLw",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "10",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book35 := models.BookModel{
		Name:                "Người Giàu Có Nhất Thành Babylon",
		Publisher:           "GEORGE SAMUEL CLASON",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/79/2000/3011.jpg?hmac=TQsXWj0kLBLRXbSAh2Pygog1-cOefqpjEoKyl0uD3tg",
		Author:              "GEORGE SAMUEL CLASON",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "10",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "GEORGE SAMUEL CLASON",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book36 := models.BookModel{
		Name:                "Storytelling With Data - Kể Chuyện Thông Qua Dữ Liệu",
		Publisher:           "Cole Nussbaumer Knaflic",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/78/1584/2376.jpg?hmac=80UVSwpa9Nfcq7sQ5kxb9Z5sD2oLsbd5zkFO5ybMC4E",
		Author:              "Cole Nussbaumer Knaflic",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "10",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Cole Nussbaumer Knaflic",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book37 := models.BookModel{
		Name:                "Basic Economics: Kinh tế học cơ bản, a bờ cờ, kinh tế học nhập môn cho nhà đầu tư",
		Publisher:           "Thomas Sowell",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/91/3504/2336.jpg?hmac=tK6z7RReLgUlCuf4flDKeg57o6CUAbgklgLsGL0UowU",
		Author:              "Thomas Sowell",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "10",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Thomas Sowell",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book38 := models.BookModel{
		Name:                "Object Oriented Programming for Java",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/92/3568/2368.jpg?hmac=k-61p7oMQe6U59dEgm0Gu6bWTJGPfHblWxMskxTBZMo",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "3",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book39 := models.BookModel{
		Name:                "Destination C1&C2 Grammar and Vocabulary",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/93/2000/1334.jpg?hmac=HdhcVTbAYkFCXsu1qBRWeEPiy05Qjc3LbnMWJlfEFjo",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "8",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book40 := models.BookModel{
		Name:                "2000 English Collocation and Idioms",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/94/2133/1200.jpg?hmac=renErwq5Y1V0wU0ah6ssb9JR6lzlhs9n05UuPzO-v0M",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "8",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book41 := models.BookModel{
		Name:                "Từ Điển Tiếng Việt (Hoàng Phê)",
		Publisher:           "Hoàng Phê",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/95/2048/2048.jpg?hmac=rvEUpzQSeSWTzfsWTIytYnLieOx_Iaa6PfYOxVwEb1g",
		Author:              "Hoàng Phê",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "9",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Hoàng Phê",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book42 := models.BookModel{
		Name:                "Từ Điển Tâm Lý - Tính Cách Và Cảm Xúc Đến Từ Đâu ?",
		Publisher:           "Shozo Shibuya",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/98/3264/2176.jpg?hmac=yRaOwMpmio9mwf43lbPEYI_5-WiPWoghJZyOKldQ43U",
		Author:              "Shozo Shibuya",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "9",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Shozo Shibuya",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book43 := models.BookModel{
		Name:                "Hán Việt Tự Điển (Vàng)",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/99/4912/3264.jpg?hmac=jobkGP8_9Sch9BmMGe3xmm8yjCVQ3iSHrbu_4kOOciY",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "9",
		Publishing_location: "VietNam",
		Page:                500,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book44 := models.BookModel{
		Name:                "Từ Điển Oxford Anh - Việt (Hơn 350.000 Từ)",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/100/2500/1656.jpg?hmac=gWyN-7ZB32rkAjMhKXQgdHOIBRHyTSgzuOK6U0vXb1w",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "9",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book45 := models.BookModel{
		Name:                "Từ Điển Anh - Anh - Việt Dành Cho Học Sinh",
		Publisher:           "TKLong",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/101/2621/1747.jpg?hmac=cu15YGotS0gIYdBbR1he5NtBLZAAY6aIY5AbORRAngs",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "9",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book46 := models.BookModel{
		Name:                "90 ngày cùng con dậy thì tỏa sáng",
		Publisher:           "OEM",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/101/2621/1747.jpg?hmac=cu15YGotS0gIYdBbR1he5NtBLZAAY6aIY5AbORRAngs",
		Author:              "Tăng Kim Long",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "9",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tăng Kim Long",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book47 := models.BookModel{
		Name:                "CHÀO TUỔI DẬY THÌ! Kiến thức về dậy thì dành cho các bạn NỮ (8-12 tuổi)",
		Publisher:           "Trần Thị Huyên Thảo",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/102/4320/3240.jpg?hmac=ico2KysoswVG8E8r550V_afIWN963F6ygTVrqHeHeRc",
		Author:              "Trần Thị Huyên Thảo",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "7",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Trần Thị Huyên Thảo",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book48 := models.BookModel{
		Name:                "Yêu Trong Tỉnh Thức - Từ Bạn Đời Đến Bạn Đạo Tập 1",
		Publisher:           "Tuệ An",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/103/2592/1936.jpg?hmac=aC1FT3vX9bCVMIT-KXjHLhP6vImAcsyGCH49vVkAjPQ",
		Author:              "Tuệ An",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "7",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Tuệ An",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book49 := models.BookModel{
		Name:                "CHÀO TUỔI DẬY THÌ! Kiến thức về dậy thì dành cho các bạn NAM (8-12 tuổi)",
		Publisher:           "Trần Thị Huyên Thảo",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/104/3840/2160.jpg?hmac=Rv0qxBiYb65Htow4mdeDlyT5kLM23Z2cDlN53YYldZU",
		Author:              "Trần Thị Huyên Thảo",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "7",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Trần Thị Huyên Thảo",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}
	Book50 := models.BookModel{
		Name:                "Sách Tâm Lý Tuổi Teen - Nhật ký tuổi teen: Mẹ hãy buông tay để con được lớn",
		Publisher:           "Đào Thiên An",
		Yearpublished:       2000,
		Book_img:            "https://fastly.picsum.photos/id/106/2592/1728.jpg?hmac=E1-3Hac5ffuCVwYwexdHImxbMFRsv83exZ2EhlYxkgY",
		Author:              "Đào Thiên An",
		Borrowed_quantity:   0,
		Amount:              0,
		Type_id:             "7",
		Publishing_location: "VietNam",
		Page:                222,
		License:             "Đào Thiên An",
		Details:             `Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis pellentesque euismod tellus pharetra placerat. Sed quis sapien nisi. Duis felis tellus, lobortis a tempor ac, accumsan ut neque. Fusce hendrerit leo non ligula convallis, ut consectetur lacus gravida. Donec eu leo ut sem vulputate mollis. Curabitur finibus ipsum.`,
		Description:         `Suspendisse ut varius libero. Curabitur consectetur fringilla ante sodales suscipit. Aenean sed neque felis. Quisque mollis diam eu odio interdum ornare. Nullam facilisis metus quam, nec ultricies mauris congue id. Donec eget lobortis velit. Etiam sit amet consectetur erat. Nunc egestas nisl quis bibendum dictum`,
	}

	Book1.Book_id = "B1"
	Book2.Book_id = "B2"
	Book3.Book_id = "B3"
	Book4.Book_id = "B4"
	Book5.Book_id = "B5"
	Book6.Book_id = "B6"
	Book7.Book_id = "B7"
	Book8.Book_id = "B8"
	Book9.Book_id = "B9"
	Book10.Book_id = "B10"
	Book11.Book_id = "B11"
	Book12.Book_id = "B12"
	Book13.Book_id = "B13"
	Book14.Book_id = "B14"
	Book15.Book_id = "B15"
	Book16.Book_id = "B16"
	Book17.Book_id = "B17"
	Book18.Book_id = "B18"
	Book19.Book_id = "B19"
	Book20.Book_id = "B20"
	Book21.Book_id = "B21"
	Book22.Book_id = "B22"
	Book23.Book_id = "B23"
	Book24.Book_id = "B24"
	Book25.Book_id = "B25"
	Book26.Book_id = "B26"
	Book27.Book_id = "B27"
	Book28.Book_id = "B28"
	Book29.Book_id = "B29"
	Book30.Book_id = "B30"
	Book31.Book_id = "B31"
	Book32.Book_id = "B32"
	Book33.Book_id = "B33"
	Book34.Book_id = "B34"
	Book35.Book_id = "B35"
	Book36.Book_id = "B36"
	Book37.Book_id = "B37"
	Book38.Book_id = "B38"
	Book39.Book_id = "B39"
	Book40.Book_id = "B40"
	Book41.Book_id = "B41"
	Book42.Book_id = "B42"
	Book43.Book_id = "B43"
	Book44.Book_id = "B44"
	Book45.Book_id = "B45"
	Book46.Book_id = "B46"
	Book47.Book_id = "B47"
	Book48.Book_id = "B48"
	Book49.Book_id = "B49"
	Book50.Book_id = "B50"
	
	Book1.Id = primitive.NewObjectID()
	Book2.Id = primitive.NewObjectID()
	Book3.Id = primitive.NewObjectID()
	Book4.Id = primitive.NewObjectID()
	Book5.Id = primitive.NewObjectID()
	Book6.Id = primitive.NewObjectID()
	Book7.Id = primitive.NewObjectID()
	Book8.Id = primitive.NewObjectID()
	Book9.Id = primitive.NewObjectID()
	Book10.Id = primitive.NewObjectID()
	Book11.Id = primitive.NewObjectID()
	Book12.Id = primitive.NewObjectID()
	Book13.Id = primitive.NewObjectID()
	Book14.Id = primitive.NewObjectID()
	Book15.Id = primitive.NewObjectID()
	Book16.Id = primitive.NewObjectID()
	Book17.Id = primitive.NewObjectID()
	Book18.Id = primitive.NewObjectID()
	Book19.Id = primitive.NewObjectID()
	Book20.Id = primitive.NewObjectID()
	Book21.Id = primitive.NewObjectID()
	Book22.Id = primitive.NewObjectID()
	Book23.Id = primitive.NewObjectID()
	Book24.Id = primitive.NewObjectID()
	Book25.Id = primitive.NewObjectID()
	Book26.Id = primitive.NewObjectID()
	Book27.Id = primitive.NewObjectID()
	Book28.Id = primitive.NewObjectID()
	Book29.Id = primitive.NewObjectID()
	Book30.Id = primitive.NewObjectID()
	Book31.Id = primitive.NewObjectID()
	Book32.Id = primitive.NewObjectID()
	Book33.Id = primitive.NewObjectID()
	Book34.Id = primitive.NewObjectID()
	Book35.Id = primitive.NewObjectID()
	Book36.Id = primitive.NewObjectID()
	Book37.Id = primitive.NewObjectID()
	Book38.Id = primitive.NewObjectID()
	Book39.Id = primitive.NewObjectID()
	Book40.Id = primitive.NewObjectID()
	Book41.Id = primitive.NewObjectID()
	Book42.Id = primitive.NewObjectID()
	Book43.Id = primitive.NewObjectID()
	Book44.Id = primitive.NewObjectID()
	Book45.Id = primitive.NewObjectID()
	Book46.Id = primitive.NewObjectID()
	Book47.Id = primitive.NewObjectID()
	Book48.Id = primitive.NewObjectID()
	Book49.Id = primitive.NewObjectID()
	Book50.Id = primitive.NewObjectID()
	Book1.Created_at = week1
	Book2.Created_at = week1
	Book3.Created_at = week1
	Book4.Created_at = week1
	Book5.Created_at = week1
	Book6.Created_at = week1
	Book7.Created_at = week1
	Book8.Created_at = week1
	Book9.Created_at = week1
	Book10.Created_at = week1
	Book11.Created_at = week1
	Book12.Created_at = week1
	Book13.Created_at = week1
	Book14.Created_at = week1
	Book15.Created_at = week1
	Book16.Created_at = week1
	Book17.Created_at = week1
	Book18.Created_at = week1
	Book19.Created_at = week1
	Book20.Created_at = week1
	Book21.Created_at = week1
	Book22.Created_at = week1
	Book23.Created_at = week1
	Book24.Created_at = week1
	Book25.Created_at = week1
	Book26.Created_at = week1
	Book27.Created_at = week1
	Book28.Created_at = week1
	Book29.Created_at = week1
	Book30.Created_at = week1
	Book31.Created_at = week1
	Book32.Created_at = week1
	Book33.Created_at = week1
	Book34.Created_at = week1
	Book35.Created_at = week1
	Book36.Created_at = week1
	Book37.Created_at = week1
	Book38.Created_at = week1
	Book39.Created_at = week1
	Book40.Created_at = week1
	Book41.Created_at = week1
	Book42.Created_at = week1
	Book43.Created_at = week1
	Book44.Created_at = week1
	Book45.Created_at = week1
	Book46.Created_at = week1
	Book47.Created_at = week1
	Book48.Created_at = week1
	Book49.Created_at = week1
	Book50.Created_at = week1


	Book1.Updated_at = week1
	Book2.Updated_at = week1
	Book3.Updated_at = week1
	Book4.Updated_at = week1
	Book5.Updated_at = now
	Book6.Updated_at = now
	Book7.Updated_at = now
	Book8.Updated_at = now
	Book9.Updated_at = week1
	Book10.Updated_at = now
	Book11.Updated_at = now
	Book12.Updated_at = week1
	Book13.Updated_at = week1
	Book14.Updated_at = week1
	Book15.Updated_at = week1
	Book16.Updated_at = week1
	Book17.Updated_at = week1
	Book18.Updated_at = week1
	Book19.Updated_at = week1
	Book20.Updated_at = week1
	Book21.Updated_at = week1
	Book22.Updated_at = week1
	Book23.Updated_at = week1
	Book24.Updated_at = week1
	Book25.Updated_at = week1
	Book26.Updated_at = week1
	Book27.Updated_at = week1
	Book28.Updated_at = week1
	Book29.Updated_at = week1
	Book30.Updated_at = week1
	Book31.Updated_at = week1
	Book32.Updated_at = week1
	Book33.Updated_at = week1
	Book34.Updated_at = week1
	Book35.Updated_at = week1
	Book36.Updated_at = week1
	Book37.Updated_at = now
	Book38.Updated_at = week1
	Book39.Updated_at = week1
	Book40.Updated_at = now
	Book41.Updated_at = week1
	Book42.Updated_at = week1
	Book43.Updated_at = now
	Book44.Updated_at = week1
	Book45.Updated_at = now
	Book46.Updated_at = week1
	Book47.Updated_at = week1
	Book48.Updated_at = week1
	Book49.Updated_at = now
	Book50.Updated_at = week1

	books := []interface{}{}
	books = append(books, Book1)
	books = append(books, Book2)
	books = append(books, Book3)
	books = append(books, Book4)
	books = append(books, Book5)
	books = append(books, Book6)
	books = append(books, Book7)
	books = append(books, Book8)
	books = append(books, Book9)
	books = append(books, Book10)
	books = append(books, Book11)
	books = append(books, Book12)
	books = append(books, Book13)
	books = append(books, Book14)
	books = append(books, Book15)
	books = append(books, Book16)
	books = append(books, Book17)
	books = append(books, Book18)
	books = append(books, Book19)
	books = append(books, Book20)
	books = append(books, Book21)
	books = append(books, Book22)
	books = append(books, Book23)
	books = append(books, Book24)
	books = append(books, Book25)
	books = append(books, Book26)
	books = append(books, Book27)
	books = append(books, Book28)
	books = append(books, Book29)
	books = append(books, Book30)
	books = append(books, Book31)
	books = append(books, Book32)
	books = append(books, Book33)
	books = append(books, Book34)
	books = append(books, Book35)
	books = append(books, Book36)
	books = append(books, Book37)
	books = append(books, Book38)
	books = append(books, Book39)
	books = append(books, Book40)
	books = append(books, Book41)
	books = append(books, Book42)
	books = append(books, Book43)
	books = append(books, Book44)
	books = append(books, Book45)
	books = append(books, Book46)
	books = append(books, Book47)
	books = append(books, Book48)
	books = append(books, Book49)
	books = append(books, Book50)
	BookCollection.InsertMany(ctx, books)
}

func SeedBookDetail() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	BookDetail1 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail2 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail3 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail4 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail5 := models.BookDetailModel{
		Book_id:    "B1",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail6 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail7 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail8 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail9 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail10 := models.BookDetailModel{
		Book_id:    "B2",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail11 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail12 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail13 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail14 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail15 := models.BookDetailModel{
		Book_id:    "B3",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail16 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail17 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail18 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail19 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail20 := models.BookDetailModel{
		Book_id:    "B4",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail21 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail22 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail23 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail24 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail25 := models.BookDetailModel{
		Book_id:    "B5",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail26 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail27 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail28 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail29 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail30 := models.BookDetailModel{
		Book_id:    "B6",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail31 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail32 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail33 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail34 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail35 := models.BookDetailModel{
		Book_id:    "B7",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail36 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail37 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail38 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail39 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail40 := models.BookDetailModel{
		Book_id:    "B8",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail41 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail42 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail43 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail44 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail45 := models.BookDetailModel{
		Book_id:    "B9",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail46 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail47 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail48 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail49 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail50 := models.BookDetailModel{
		Book_id:    "B10",
		Location:   "ABC123",
		Status:     "ready",
		Created_at: now,
		Updated_at: now,
	}
	BookDetail1.Id = primitive.NewObjectID()
	BookDetail2.Id = primitive.NewObjectID()
	BookDetail3.Id = primitive.NewObjectID()
	BookDetail4.Id = primitive.NewObjectID()
	BookDetail5.Id = primitive.NewObjectID()
	BookDetail6.Id = primitive.NewObjectID()
	BookDetail7.Id = primitive.NewObjectID()
	BookDetail8.Id = primitive.NewObjectID()
	BookDetail9.Id = primitive.NewObjectID()
	BookDetail10.Id = primitive.NewObjectID()
	BookDetail11.Id = primitive.NewObjectID()
	BookDetail12.Id = primitive.NewObjectID()
	BookDetail13.Id = primitive.NewObjectID()
	BookDetail14.Id = primitive.NewObjectID()
	BookDetail15.Id = primitive.NewObjectID()
	BookDetail16.Id = primitive.NewObjectID()
	BookDetail17.Id = primitive.NewObjectID()
	BookDetail18.Id = primitive.NewObjectID()
	BookDetail19.Id = primitive.NewObjectID()
	BookDetail20.Id = primitive.NewObjectID()
	BookDetail21.Id = primitive.NewObjectID()
	BookDetail22.Id = primitive.NewObjectID()
	BookDetail23.Id = primitive.NewObjectID()
	BookDetail24.Id = primitive.NewObjectID()
	BookDetail25.Id = primitive.NewObjectID()
	BookDetail25.Id = primitive.NewObjectID()
	BookDetail26.Id = primitive.NewObjectID()
	BookDetail27.Id = primitive.NewObjectID()
	BookDetail28.Id = primitive.NewObjectID()
	BookDetail29.Id = primitive.NewObjectID()
	BookDetail30.Id = primitive.NewObjectID()
	BookDetail31.Id = primitive.NewObjectID()
	BookDetail32.Id = primitive.NewObjectID()
	BookDetail33.Id = primitive.NewObjectID()
	BookDetail34.Id = primitive.NewObjectID()
	BookDetail35.Id = primitive.NewObjectID()
	BookDetail36.Id = primitive.NewObjectID()
	BookDetail37.Id = primitive.NewObjectID()
	BookDetail38.Id = primitive.NewObjectID()
	BookDetail39.Id = primitive.NewObjectID()
	BookDetail40.Id = primitive.NewObjectID()
	BookDetail41.Id = primitive.NewObjectID()
	BookDetail42.Id = primitive.NewObjectID()
	BookDetail43.Id = primitive.NewObjectID()
	BookDetail44.Id = primitive.NewObjectID()
	BookDetail45.Id = primitive.NewObjectID()
	BookDetail46.Id = primitive.NewObjectID()
	BookDetail47.Id = primitive.NewObjectID()
	BookDetail48.Id = primitive.NewObjectID()
	BookDetail49.Id = primitive.NewObjectID()
	BookDetail50.Id = primitive.NewObjectID()

	BookDetail1.Book_detail_id = "D11"
	BookDetail2.Book_detail_id = "D12"
	BookDetail3.Book_detail_id = "D13"
	BookDetail4.Book_detail_id = "D14"
	BookDetail5.Book_detail_id = "D15"
	BookDetail6.Book_detail_id = "D21"
	BookDetail7.Book_detail_id = "D22"
	BookDetail8.Book_detail_id = "D23"
	BookDetail9.Book_detail_id = "D24"
	BookDetail10.Book_detail_id = "D25"
	BookDetail11.Book_detail_id = "D31"
	BookDetail12.Book_detail_id = "D32"
	BookDetail13.Book_detail_id = "D33"
	BookDetail14.Book_detail_id = "D34"
	BookDetail15.Book_detail_id = "D35"
	BookDetail16.Book_detail_id = "D41"
	BookDetail17.Book_detail_id = "D42"
	BookDetail18.Book_detail_id = "D43"
	BookDetail19.Book_detail_id = "D44"
	BookDetail20.Book_detail_id = "D45"
	BookDetail21.Book_detail_id = "D51"
	BookDetail22.Book_detail_id = "D52"
	BookDetail23.Book_detail_id = "D53"
	BookDetail24.Book_detail_id = "D54"
	BookDetail25.Book_detail_id = "D55"
	BookDetail26.Book_detail_id = "D61"
	BookDetail27.Book_detail_id = "D62"
	BookDetail28.Book_detail_id = "D63"
	BookDetail29.Book_detail_id = "D64"
	BookDetail30.Book_detail_id = "D65"
	BookDetail31.Book_detail_id = "D71"
	BookDetail32.Book_detail_id = "D72"
	BookDetail33.Book_detail_id = "D73"
	BookDetail34.Book_detail_id = "D74"
	BookDetail35.Book_detail_id = "D75"
	BookDetail36.Book_detail_id = "D81"
	BookDetail37.Book_detail_id = "D82"
	BookDetail38.Book_detail_id = "D83"
	BookDetail39.Book_detail_id = "D84"
	BookDetail40.Book_detail_id = "D85"
	BookDetail41.Book_detail_id = "D91"
	BookDetail42.Book_detail_id = "D92"
	BookDetail43.Book_detail_id = "D93"
	BookDetail44.Book_detail_id = "D94"
	BookDetail45.Book_detail_id = "D95"
	BookDetail46.Book_detail_id = "D101"
	BookDetail47.Book_detail_id = "D102"
	BookDetail48.Book_detail_id = "D103"
	BookDetail49.Book_detail_id = "D104"
	BookDetail50.Book_detail_id = "D105"

	BookDetail := []interface{}{}
	BookDetail = append(BookDetail, BookDetail1)
	BookDetail = append(BookDetail, BookDetail2)
	BookDetail = append(BookDetail, BookDetail3)
	BookDetail = append(BookDetail, BookDetail4)
	BookDetail = append(BookDetail, BookDetail5)
	BookDetail = append(BookDetail, BookDetail6)
	BookDetail = append(BookDetail, BookDetail7)
	BookDetail = append(BookDetail, BookDetail8)
	BookDetail = append(BookDetail, BookDetail9)
	BookDetail = append(BookDetail, BookDetail10)
	BookDetail = append(BookDetail, BookDetail11)
	BookDetail = append(BookDetail, BookDetail12)
	BookDetail = append(BookDetail, BookDetail13)
	BookDetail = append(BookDetail, BookDetail14)
	BookDetail = append(BookDetail, BookDetail15)
	BookDetail = append(BookDetail, BookDetail16)
	BookDetail = append(BookDetail, BookDetail17)
	BookDetail = append(BookDetail, BookDetail18)
	BookDetail = append(BookDetail, BookDetail19)
	BookDetail = append(BookDetail, BookDetail20)
	BookDetail = append(BookDetail, BookDetail21)
	BookDetail = append(BookDetail, BookDetail22)
	BookDetail = append(BookDetail, BookDetail23)
	BookDetail = append(BookDetail, BookDetail24)
	BookDetail = append(BookDetail, BookDetail25)
	BookDetail = append(BookDetail, BookDetail26)
	BookDetail = append(BookDetail, BookDetail27)
	BookDetail = append(BookDetail, BookDetail28)
	BookDetail = append(BookDetail, BookDetail29)
	BookDetail = append(BookDetail, BookDetail30)
	BookDetail = append(BookDetail, BookDetail31)
	BookDetail = append(BookDetail, BookDetail32)
	BookDetail = append(BookDetail, BookDetail33)
	BookDetail = append(BookDetail, BookDetail34)
	BookDetail = append(BookDetail, BookDetail35)
	BookDetail = append(BookDetail, BookDetail36)
	BookDetail = append(BookDetail, BookDetail37)
	BookDetail = append(BookDetail, BookDetail38)
	BookDetail = append(BookDetail, BookDetail39)
	BookDetail = append(BookDetail, BookDetail40)
	BookDetail = append(BookDetail, BookDetail41)
	BookDetail = append(BookDetail, BookDetail42)
	BookDetail = append(BookDetail, BookDetail43)
	BookDetail = append(BookDetail, BookDetail44)
	BookDetail = append(BookDetail, BookDetail45)
	BookDetail = append(BookDetail, BookDetail46)
	BookDetail = append(BookDetail, BookDetail47)
	BookDetail = append(BookDetail, BookDetail48)
	BookDetail = append(BookDetail, BookDetail49)
	BookDetail = append(BookDetail, BookDetail50)
	BookDetailCollection.InsertMany(ctx, BookDetail)
}

func SeedBookRent() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	BookRent1 := models.BookRentModel{
		Book_id: "nil",
		User_id: "nil",
	}
	// BookRent2 := models.BookRentModel{
	// 	Book_id:      "B5",
	// 	User_id:      "U03",
	// 	Reserve_date: now,
	// }
	// BookRent3 := models.BookRentModel{
	// 	Book_id:      "B5",
	// 	User_id:      "U04",
	// 	Reserve_date: now,
	// }
	// BookRent4 := models.BookRentModel{
	// 	Book_id:      "B10",
	// 	User_id:      "U05",
	// 	Reserve_date: now,
	// }
	BookRent1.Book_detail_id = "nil"
	// BookRent2.Book_detail_id = "D51"
	// BookRent3.Book_detail_id = "D52"
	// BookRent4.Book_detail_id = "D105"
	BookRent1.Book_rent_id = "nil"
	// BookRent2.Book_rent_id = "R2"
	// BookRent3.Book_rent_id = "R3"
	// BookRent4.Book_rent_id = "R4"

	BookRent1.Id = primitive.NewObjectID()
	// BookRent2.Id = primitive.NewObjectID()
	// BookRent3.Id = primitive.NewObjectID()
	// BookRent4.Id = primitive.NewObjectID()

	// BookRent := []interface{}{
	// 	BookRent1,
	// 	BookRent2,
	// 	BookRent3,
	// 	BookRent4,
	// }
	BookRentCollection.InsertOne(ctx, BookRent1)
}

func SeedBookBorrow() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	BookBorrow1 := models.BookBorrowModel{
		User_id:        "nil",
		Book_id:        "nil",
		Book_detail_id: "nil",
	}
	// BookBorrow2 := models.BookBorrowModel{
	// 	User_id:        "U03",
	// 	Book_id:        "B6",
	// 	Book_detail_id: "D62",
	// }
	// BookBorrow3 := models.BookBorrowModel{
	// 	User_id:        "U04",
	// 	Book_id:        "B7",
	// 	Book_detail_id: "D75",
	// }
	// BookBorrow4 := models.BookBorrowModel{
	// 	User_id:        "U05",
	// 	Book_id:        "B8",
	// 	Book_detail_id: "D81",
	// }
	BookBorrow1.Id = primitive.NewObjectID()
	// BookBorrow2.Id = primitive.NewObjectID()
	// BookBorrow3.Id = primitive.NewObjectID()
	// BookBorrow4.Id = primitive.NewObjectID()
	BookBorrow1.Book_hire_id = "nil"
	// BookBorrow2.Book_hire_id = "BR02"
	// BookBorrow3.Book_hire_id = "BR03"
	// BookBorrow4.Book_hire_id = "BR04"
	// BookBorrows := []interface{}{
	// 	BookBorrow1,
	// 	BookBorrow2,
	// 	BookBorrow3,
	// 	BookBorrow4,
	// }
	BookBorrowCollection.InsertOne(ctx, BookBorrow1)
}

func SeedHistory() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	History1 := models.HistoryModel{
		User_id:        "U02",
		Book_id:        "B3",
		Book_detail_id: "D31",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History2 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B5",
		Book_detail_id: "D51",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History3 := models.HistoryModel{
		User_id:        "U04",
		Book_id:        "B10",
		Book_detail_id: "D101",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History4 := models.HistoryModel{
		User_id:        "U05",
		Book_id:        "B10",
		Book_detail_id: "D101",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History5 := models.HistoryModel{
		User_id:        "U02",
		Book_id:        "B10",
		Book_detail_id: "D102",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History6 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B10",
		Book_detail_id: "D103",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History7 := models.HistoryModel{
		User_id:        "U04",
		Book_id:        "B10",
		Book_detail_id: "D104",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History8 := models.HistoryModel{
		User_id:        "U02",
		Book_id:        "B4",
		Book_detail_id: "D41",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History9 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B5",
		Book_detail_id: "D53",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History10 := models.HistoryModel{
		User_id:        "U04",
		Book_id:        "B6",
		Book_detail_id: "D64",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History11 := models.HistoryModel{
		User_id:        "U05",
		Book_id:        "B7",
		Book_detail_id: "D71",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History12 := models.HistoryModel{
		User_id:        "U03",
		Book_id:        "B9",
		Book_detail_id: "D93",
		Status:         "returned",
		Date_borrowed:  week1,
		Date_return:    now,
	}
	History1.Id = primitive.NewObjectID()
	History2.Id = primitive.NewObjectID()
	History3.Id = primitive.NewObjectID()
	History4.Id = primitive.NewObjectID()
	History5.Id = primitive.NewObjectID()
	History6.Id = primitive.NewObjectID()
	History7.Id = primitive.NewObjectID()
	History8.Id = primitive.NewObjectID()
	History9.Id = primitive.NewObjectID()
	History10.Id = primitive.NewObjectID()
	History11.Id = primitive.NewObjectID()
	History12.Id = primitive.NewObjectID()
	History1.History_id = "H01"
	History2.History_id = "H02"
	History3.History_id = "H03"
	History4.History_id = "H04"
	History5.History_id = "H05"
	History6.History_id = "H06"
	History7.History_id = "H07"
	History8.History_id = "H08"
	History9.History_id = "H09"
	History10.History_id = "H10"
	History11.History_id = "H11"
	History12.History_id = "H12"
	Histories := []interface{}{
		History1,
		History2,
		History3,
		History4,
		History5,
		History6,
		History7,
		History8,
		History9,
		History10,
		History11,
		History12,
	}
	HistoryCollection.InsertMany(ctx, Histories)
}

func SeedType() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	Type1 := models.BookTypeModel{
		TypeId:   "1",
		TypeName: "Manga",
	}
	Type2 := models.BookTypeModel{
		TypeId:   "2",
		TypeName: "Education",
	}
	Type3 := models.BookTypeModel{
		TypeId:   "3",
		TypeName: "IT",
	}
	Type4 := models.BookTypeModel{
		TypeId:   "4",
		TypeName: "Tôn Giáo - Tâm Linh",
	}
	Type5 := models.BookTypeModel{
		TypeId:   "5",
		TypeName: "Tiểu Thuyết",
	}
	Type6 := models.BookTypeModel{
		TypeId:   "6",
		TypeName: "Light Novel",
	}
	Type7 := models.BookTypeModel{
		TypeId:   "7",
		TypeName: "Tâm Lý",
	}
	Type8 := models.BookTypeModel{
		TypeId:   "8",
		TypeName: "Học Ngoại Ngữ",
	}
	Type9 := models.BookTypeModel{
		TypeId:   "9",
		TypeName: "Từ Điển",
	}
	Type10 := models.BookTypeModel{
		TypeId:   "10",
		TypeName: "Kinh tế",
	}
	Type1.Id = primitive.NewObjectID()
	Type2.Id = primitive.NewObjectID()
	Type3.Id = primitive.NewObjectID()
	Type4.Id = primitive.NewObjectID()
	Type5.Id = primitive.NewObjectID()
	Type6.Id = primitive.NewObjectID()
	Type7.Id = primitive.NewObjectID()
	Type8.Id = primitive.NewObjectID()
	Type9.Id = primitive.NewObjectID()
	Type10.Id = primitive.NewObjectID()
	Types := []interface{}{
		Type1,
		Type2,
		Type3,
		Type4,
		Type5,
		Type6,
		Type7,
		Type8,
		Type9,
		Type10,
	}
	TypeCollections.InsertMany(ctx, Types)
}
