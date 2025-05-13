## 🧩 **1. Hiện trạng – Vấn đề đang gặp phải**
### 🛠 **1.1. Phát triển tính năng rời rạc**
-   **Vấn đề**: Các tính năng phát triển trong thời gian dài riêng rẽ, gây ra các vấn đề:
    -   Xung đột logic/mã nguồn (source code)
    -   Cùng thực hiện một công việc mà không biết

-   **Ví dụ tương đồng trong thực tế**:  
    Giả sử công ty muốn cải tiến giao diện trang báo cáo tài sản cho nhà đầu tư.
    -   Nhóm **Phát triển sản phẩm** đưa ra đề xuất thiết kế mới.
    -   Nhóm **IT** song song cũng đang tự triển khai một phiên bản nâng cấp để xử lý phản hồi cũ của khách hàng.
    -   Hai nhóm đều không biết về công việc của nhau, không trao đổi định kỳ.
    
    👉 Đến khi gần ra mắt:
    -   Cả hai nhóm đều làm trùng phần giao diện, nhưng logic xử lý lại khác nhau.
    -   Cần tốn thêm người để rà soát, hợp nhất, kiểm thử lại từ đầu.
    -   Dẫn đến **lãng phí công sức** cả tháng trời và chậm ra mắt sản phẩm.
### ⏳ **1.2. Tốn thời gian cho việc không tạo ra giá trị thực (Development Overhead)**

-   **Vấn đề**: Trong quá trình phát triển phần mềm, một phần đáng kể thời gian của team bị lãng phí vào các công việc “phụ” như:
    -   Chờ phần khác hoàn thành để tiếp tục
    -   Làm rõ yêu cầu/chi tiết với các bên
    -   Xử lí các lỗi xung đột do tích hợp nhiều tính năng phát triển trong thời gian dài
    -   Làm lại các công việc mà trước đây đi sai hướng
    -   Do tính chất số lượng yêu cầu nhiều, thay đổi thường xuyên, nên tốn thời gian replan công việc

-   **Ảnh hưởng cụ thể**:
	-   Dự án bị kéo dài không phải do độ khó, mà do cách phối hợp chưa hiệu quả.
	-   Nhân viên kỹ thuật làm việc mệt mỏi, không thấy “kết quả” rõ ràng.
	-   Quản lý thì khó đánh giá năng suất thật vì không thấy được các công việc phụ này.

### 🧯 **1.3. Yêu cầu hoàn thành 100% trước khi golive – Thiếu linh hoạt trong môi trường nhiều biến động**

-   **Vấn đề**: Quy trình hiện tại yêu cầu tất cả tính năng trong một sprint phải hoàn thiện hoàn toàn trước khi được golive. Nếu chỉ một phần gặp vấn đề, **toàn bộ phiên bản bị trì hoãn**, hoặc **tính năng chưa hoàn thiện bị loại bỏ**, dẫn đến:
    -   Một số tính năng đã tiêu tốn công sức nhưng không đem lại giá trị.
    -   Tạo áp lực phải hoàn thành tính năng đúng hạn.
-   **Ví dụ thực tế**:  UBCK thay đổi luật, yêu cầu thay đổi tính năng hiện có
    -   Team thực hiện đánh giá lại sprint, loại bỏ tính năng khác để ưu tiên đáp ứng nhu cầu pháp lí
    -   Dev thực hiện loại bỏ tính năng ra khỏi nhánh tích hợp trước khi golive
    -   QA thực hiện verify tính năng được loại bỏ

### 🧯 **1.4. Các tác động hệ thống diễn ra muộn**

-   **Vấn đề**: Do đặc thù về nghiệp vụ, các thay đổi đối với hệ thống thực cần thực hiện ngoài giờ giao dịch để đảm bảo an toàn, dẫn đến:
    -   Với những lần nâng cấp lớn, nhiều tính năng, team thường phải làm muộn.
    -   Có những tính năng phải chuyển sang ngày nghỉ.
    -   Team phát triển ít thời gian dành cho cá nhân.
-   **Ví dụ thực tế**:  UBCK thay đổi luật, yêu cầu thay đổi tính năng hiện có
    -   Team thực hiện đánh giá lại sprint, loại bỏ tính năng khác để ưu tiên đáp ứng nhu cầu pháp lí
    -   Dev thực hiện loại bỏ tính năng ra khỏi nhánh tích hợp trước khi golive
    -   QA thực hiện verify tính năng được loại bỏ
👉 Tiêu tốn nhiều thời gian của team phát triển, dẫn đến thiếu thời gian thực hiện cập nhật theo luật, gây ra áp lực phải hoàn thành cho team

## 🔧 **2. Giải pháp cải tiến quy trình**

### 🌱 **2.1. Trunk-Based Development – Làm việc cùng nhau, không tách rời**

-   **Giải thích ngắn gọn**:  
    Thay vì mỗi người làm một phần riêng biệt rồi gộp lại sau, cả nhóm cùng làm việc trên một “nhánh chính” duy nhất, cập nhật thường xuyên theo ngày.
    
-   **Ví dụ thực tế**:  
    Giống như việc nhiều nhóm cùng viết một văn bản Google Docs – mọi thay đổi được cập nhật liên tục, tránh trùng lặp, sai lệch.
    
-   **Giải quyết vấn đề nào?**
    
    -   Giảm trùng việc, xung đột khi tích hợp
    -   Giữ đồng bộ thông tin giữa các nhóm
    -   Tăng tính chủ động, không chờ đợi
        

----------

### 🔘 **2.2. Feature Flag – Bật/Tắt tính năng linh hoạt như công tắc**

-   **Giải thích ngắn gọn**:  
    Feature Flag cho phép lập trình viên **bật hoặc tắt từng tính năng riêng biệt**, ngay cả khi nó đã được đưa lên hệ thống. Tính năng có thể tồn tại nhưng **ẩn với người dùng** cho đến khi sẵn sàng.
    
-   **Ví dụ thực tế**:  
    Như khi bạn chuẩn bị khai trương một khu vực mới trong chi nhánh công ty, bạn có thể hoàn thiện sớm, nhưng chỉ mở cửa khi thời điểm phù hợp — không cần đợi toàn bộ mặt bằng sẵn sàng.
    
-   **Giải quyết vấn đề nào?**
    -   Không phải chờ hoàn thành 100% để golive
    -   Linh hoạt khi có thay đổi đột xuất (pháp lý, thị trường)
    -   Tránh bỏ phí tính năng đã hoàn thiện
    -   Giảm áp lực golive ngoài giờ

### 🌿 **2.3. Tư duy "đi từng bước nhỏ nhưng chắc chắn" – Cách làm hiệu quả hơn trong môi trường nhiều biến động**

-   Trong thực tế, **cố gắng đi những bước dài** để về đích nhanh thường **tiềm ẩn nhiều rủi ro**:
    -   Nếu bước sai, phải quay lại rất xa.
    -   Nếu môi trường thay đổi, hướng đi ban đầu không còn phù hợp.
        

- **Ngược lại**, việc **đi từng bước nhỏ nhưng đều đặn, có kiểm soát**, giúp chúng ta:

	-   Phát hiện sai lệch sớm
	-   Điều chỉnh dễ dàng
	-   Vẫn tiến lên mỗi ngày mà không bị gián đoạn
    

----------

### 📌 **2.4. Liên hệ với trunk-based development & feature flag:**

#### 🔧 **Trunk-Based Development**

-   Cho phép **nhóm phát triển cập nhật liên tục**, thay vì để mỗi người đi “xa” rồi quay lại ghép vào.
    
-   Giống như việc **mọi người cùng viết trên một file Google Docs**:
    
    -   Ai cũng **thấy được thay đổi ngay lập tức**
        
    -   Cùng nhìn một bản “sản phẩm cuối cùng” đang hình thành
        
    -   Tránh việc mỗi người viết một kiểu rồi đến cuối mới biết “không khớp nhau”
        

> 👉 Điều này giúp **team liên tục kiểm tra hướng đi**, điều chỉnh sớm, giảm nguy cơ sai lệch hoặc trùng lặp công việc.

----------

#### 🟢 **Feature Flags**

-   Cho phép **kích hoạt dần dần các tính năng**, thay vì đợi hoàn thiện tất cả rồi mới tung ra.
    
-   Giống như việc bạn **xây xong từng phòng trong một ngôi nhà, rồi mở từng phòng khi sẵn sàng** — không cần đợi toàn bộ căn nhà hoàn chỉnh mới cho khách vào ở.
    

> 👉 Giúp tận dụng tối đa công sức đã bỏ ra, phản hồi sớm từ người dùng, và giảm rủi ro khi cần thay đổi.
## ✅ **3. Lợi ích cụ thể khi áp dụng feature flag & trunk-based development**

----------

### **1. Giảm effort không cần thiết, tăng tập trung cho công việc tạo giá trị**
-   **Vấn đề cũ**: Developer thường phải tốn thời gian sửa lỗi merge, đợi tích hợp, hoặc xử lý những tính năng sau cùng bị drop.
-   **Giải pháp mang lại**: Làm việc trên một nhánh chính, tính năng chưa xong thì “ẩn” – không gây ảnh hưởng tới phần khác.
    
-   **Kết quả**:
    -   Dev **giảm thời gian xử lý việc thừa**, tập trung vào giá trị chính.
    -   QA **không phải test lại toàn bộ hệ thống** mỗi khi có thay đổi nhỏ, nhờ feature flags cô lập rõ phạm vi ảnh hưởng.
    

----------

### **2. Tăng tốc độ ra mắt tính năng mới**

-   **Vấn đề cũ**: Phải chờ tất cả tính năng xong mới được golive.
-   **Giải pháp mang lại**: Có thể triển khai phần đã xong, bật/tắt tính năng tùy thời điểm.
-   **Kết quả**: Tính năng tốt đến tay người dùng sớm hơn, bắt kịp thị trường.
    

----------

### **3. Thích ứng nhanh với thay đổi**

-   **Vấn đề cũ**: Gặp thay đổi pháp lý bất ngờ → dừng toàn bộ sprint, làm lại kế hoạch.
-   **Giải pháp mang lại**: Chuyển hướng linh hoạt mà không ảnh hưởng đến những phần đã làm xong.
-   **Kết quả**: Đội ngũ không bị “vỡ kế hoạch” khi có tình huống phát sinh.
    

----------

### **4. Giảm áp lực triển khai ngoài giờ**

-   **Vấn đề cũ**: Golive thường vào buổi tối/cuối tuần để tránh rủi ro.
-   **Giải pháp mang lại**: Có thể deploy trước, bật/tắt tính năng vào giờ hành chính.
-   **Kết quả**: Cải thiện work-life balance, giữ sức bền đội ngũ lâu dài.
    

----------



### **5. Tăng chất lượng sản phẩm tổng thể**
-   **Vấn đề cũ**: Lỗi phát sinh do các vấn đề tích hợp, mất thời gian xử lí và có khả năng gây ảnh hưởng đến kế hoạch.
-   **Giải pháp mang lại**: Phát hiện lỗi sớm khi tất cả cùng làm trên một nhánh chính.
-   **Kết quả**: Sản phẩm ổn định hơn, dễ bảo trì, dễ mở rộng. QA dễ dàng xác định phạm vi ảnh hưởng để kiểm thử.
----------
### **6. Trao quyền chủ động cho các bộ phận kinh doanh, giảm phụ thuộc vào team phát triển**

-   **Vấn đề cũ**:  
    Mỗi khi muốn bật/tắt một tính năng, chạy một chiến dịch thử nghiệm, hay dừng tạm một phần sản phẩm — các bộ phận nghiệp vụ như marketing, sản phẩm hay hỗ trợ khách hàng đều phải **đợi dev xử lý**, dù thay đổi đó rất nhỏ.
    
-   **Giải pháp từ Feature Flags**:  
    Giúp cấu hình bật/tắt tính năng trở thành **một hành động có thể thao tác được bởi stakeholder**, không cần sửa code hay deploy lại.
    
-   **Kết quả**:
    -   Bộ phận kinh doanh có thể chủ động **ra quyết định vào thời điểm phù hợp** (theo nhu cầu thị trường, hiệu suất hệ thống, hành vi người dùng).
    -   Team phát triển **tập trung vào cải tiến cốt lõi**, không bị phân tán vào những thay đổi nhỏ, khẩn cấp nhưng đơn giản.
### 
#### **7. Tăng độ chính xác của phản hồi – Cải tiến sản phẩm dựa trên feedback thật**

-   **Vấn đề cũ**: Feedback từ stakeholder hoặc từ môi trường staging thường không phản ánh đầy đủ trải nghiệm người dùng thật sự — do dữ liệu, hành vi, tốc độ hệ thống đều khác biệt.
    
-   **Giải pháp**: Nhờ feature flags, có thể **thử nghiệm một tính năng ngay trên production** nhưng giới hạn người dùng (nội bộ, beta user...), đảm bảo an toàn và kiểm soát.
    
-   **Kết quả**:
    
    -   Thu được phản hồi sát thực tế hơn (về cả tính năng và hiệu năng – functional & non-functional)
        
    -   Tăng tính chính xác của feedback loop
        
    -   Rút ngắn chu kỳ cải tiến sản phẩm

## 🔮 **4. Những cơ hội mở ra khi áp dụng feature flags & trunk-based development**

----------

### **1. Chủ động kinh doanh – Không còn phụ thuộc 100% vào đội phát triển**

-   **Hiện tại**: Mọi thay đổi liên quan đến sản phẩm – dù nhỏ – đều cần đến đội dev để chỉnh sửa, test, deploy.
    
-   **Cơ hội mới**: Với Feature Flags, các stakeholder (sản phẩm, kinh doanh, vận hành) có thể **tự kích hoạt hoặc tắt tính năng** vào thời điểm phù hợp, mà **không cần dev can thiệp**.
    
-   **Ý nghĩa**:
    
    -   Ra quyết định theo đúng thời điểm thị trường
        
    -   Thử nghiệm linh hoạt mà không làm gián đoạn hệ thống
        
    -   Dev tập trung vào phát triển chiều sâu, thay vì xử lý “việc vặt”
        

----------

### **2. Dễ dàng triển khai A/B testing hoặc thử nghiệm theo phân khúc người dùng**

-   **Cơ hội**: Có thể bật tính năng cho 5%, 10%, hoặc một nhóm khách hàng cụ thể để kiểm tra phản ứng thực tế.
    
-   **Lợi ích**:
    
    -   Ra quyết định dựa trên dữ liệu (data-driven)
        
    -   Tránh rủi ro từ việc rollout ồ ạt
        
    -   Phù hợp với môi trường tài chính – nơi cần cẩn trọng và xác thực trước khi mở rộng
        

----------

### **3. Chuẩn hóa quy trình CI/CD hiện đại**

-   **Feature Flags** và **Trunk-Based Development** là hai mảnh ghép thiết yếu của hệ thống phát triển hiện đại:
    
    -   Phát triển liên tục (Continuous Integration)
        
    -   Triển khai liên tục (Continuous Deployment)
        
-   **Kết quả**:
    
    -   Giảm tối đa khoảng cách giữa "code xong" và "đến tay người dùng"
        
    -   Tạo nền tảng bền vững để scale team và tăng tốc sản phẩm
        

----------

### **4. Cá nhân hóa trải nghiệm người dùng**

-   **Cơ hội**: Có thể bật những tính năng riêng cho từng nhóm khách hàng: khách hàng VIP, nhà đầu tư mới, hoặc người dùng thử nghiệm.
    
-   **Lợi ích**:
    
    -   Tăng sự hài lòng của khách hàng
        
    -   Cung cấp trải nghiệm phù hợp mà không cần tạo nhiều phiên bản hệ thống riêng biệt
        

----------

### **5. Tăng khả năng kiểm soát rủi ro – rollback không cần rollback code**

-   Nếu một tính năng gặp sự cố, không cần deploy lại phiên bản cũ — chỉ cần **tắt flag** là hệ thống trở lại trạng thái ổn định.
    
-   Giảm áp lực khi triển khai tính năng mới
    
-   Tăng độ tin cậy trong mắt khách hàng, đặc biệt trong ngành tài chính – nơi ổn định là yếu tố sống còn
----------


### **6. Tạo điều kiện để đổi mới mạnh mẽ mà vẫn giữ trải nghiệm khách hàng ổn định**

-   **Về phía đội phát triển**:
    
    -   Với feature flags, team có thể **tự tin thử nghiệm công nghệ mới, kiến trúc mới hoặc phương pháp xử lý mới** ngay trong môi trường thật.
        
    -   Có thể **triển khai dần**, chọn lọc đối tượng sử dụng, thu thập phản hồi, và cải tiến liên tục.
        
    -   **Tư duy đổi mới không còn bị cản trở bởi nỗi lo “nếu sai thì hỏng cả hệ thống”**.
        
-   **Về phía khách hàng**:
    
    -   Dù bên trong đang thử nghiệm hay thay đổi, **trải nghiệm mà khách hàng nhận được vẫn ổn định, mượt mà, nhất quán**.
        
    -   Điều này giúp **giảm thiểu lo lắng, tăng sự tin tưởng** – đặc biệt quan trọng trong lĩnh vực tài chính, nơi “ổn định” là một yêu cầu ngầm định.

## 📖 **6. Câu chuyện thực tế – Chuyển đổi hệ thống với Sở Giao dịch, tháng 5/2025**

> Đầu tháng 5 vừa rồi, team của tôi đã phải thực hiện **việc chuyển đổi hệ thống kết nối với Sở giao dịch chứng khoán** – một việc lẽ ra đã hoàn thành từ năm ngoái, nhưng bị hoãn.
> 
> Khi Sở bất ngờ kích hoạt việc chuyển đổi, chúng tôi buộc phải thực hiện lại toàn bộ quy trình trong thời gian rất ngắn.
> 
> Điều khó khăn nhất là:
> 
> -   **Hệ thống cũ và mới phải luân phiên bật/tắt để xử lý lệnh**, tạo ra rất nhiều bất tiện cho các bộ phận liên quan.
>     
> -   Chúng tôi đã phải **nâng cấp rồi rollback hệ thống**, mỗi lần gồm rất nhiều bước — tổng cộng mất **2 ngày làm việc căng thẳng**, bao gồm dev, QA, vận hành và hỗ trợ khách hàng.
>     

----------

### 💡 **Nếu đã có Feature Flags & Trunk-Based Development từ trước thì sao?**

> Nếu hệ thống đã áp dụng **feature flags**, chúng tôi hoàn toàn có thể:
> 
> -   Triển khai trước các thay đổi từ nhiều tuần trước,
>     
> -   **Ẩn toàn bộ chức năng kết nối mới dưới một flag**, không ảnh hưởng đến người dùng,
>     
> -   Và **chỉ bật công tắc** ngay khi Sở yêu cầu chuyển đổi – không cần rollback, không phải sửa gấp.
>     

> Nếu dùng **trunk-based development**, toàn bộ các thay đổi đều đã được tích hợp và kiểm thử liên tục — không có chuyện “gộp gấp” vào phút chót.

----------

### ✅ **Kết quả sẽ là:**

-   Không cần downtime kéo dài
    
-   Không gián đoạn hệ thống cho người dùng
    
-   Giảm áp lực đội kỹ thuật
    
-   Và toàn bộ quy trình **chuyển đổi sẽ diễn ra êm, gọn và không ồn ào**

## ❓ **7. Những băn khoăn từ phía quản lí**

### 1. **"Có ảnh hưởng đến tiến độ hiện tại không?"**

> → Giải thích: Không thay đổi ồ ạt. Có thể bắt đầu ở quy mô nhỏ (1 team, 1 module), song song với quy trình cũ để đánh giá hiệu quả thực tế.

----------

### 2. **"Có làm phức tạp thêm quy trình vận hành không?"**

> → Giải thích: Mục tiêu là _giảm phức tạp_, vì thay đổi sẽ rõ ràng hơn, kiểm soát dễ hơn. Stakeholder có thể thao tác mà không cần quy trình kỹ thuật dài dòng.

----------

### 3. **"Có rủi ro gì với hệ thống đang hoạt động không?"**

> → Giải thích: Chính feature flags là công cụ **giảm thiểu rủi ro** — có thể tắt tính năng ngay nếu có sự cố, không ảnh hưởng đến phần còn lại.

----------

### 4. **"Việc này có cần đầu tư công nghệ hay công cụ mới không?"**

> → Giải thích: Phần lớn có thể triển khai bằng công cụ nội bộ (open-source hoặc tích hợp vào CI/CD pipeline sẵn có). Chi phí ban đầu thấp, lợi ích lâu dài lớn.

----------

### 5. **"Có cần đào tạo lại hay thay đổi cách làm việc của team không?"**

> → Giải thích: Có điều chỉnh, nhưng không phải “đập đi làm lại”. Đội ngũ kỹ thuật chủ động chuyển đổi và lan tỏa kinh nghiệm dần dần.
