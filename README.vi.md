# tcgen: test case generator for coding problems

Tạo test case cho câu hỏi lập trình.

## Tải về

https://github.com/hoangvvo/tcgen/releases

## Hướng dẫn

https://user-images.githubusercontent.com/40987398/154596409-4c8c2dd9-1f42-4b72-851c-22b8812c56c3.mp4

Tạo 2 file code cho các mục đích:

- **Tạo case**: Khi chạy, code này sẽ viết ra file input.
- **Giải case**: Khi chạy, code này sẽ đọc file input đuợc tạo trên và viết ra file output

Các file input và output phải có đuôi đuợc liệt kê trong thiết lập `testcase.extensions`.

Chạy chuơng trình. Sau đó nhập địa điểm file Tạo case, file Giải case, và số lần chạy `n`.

Sau khi chạy `n` lần, chuơng trình sẽ copy hết các file input, output cho vào các thứ mục có tên dạng `testcase.output`.

## File thiết lập

File thiết lập tùy chọn có thể được đặt trong cùng thư mục với chuơng trình `tcgen`. Xem file thiết lập mẫu: [tcgen.yml](./gen/tcgen.yml).

- `testcase.extensions`: file có các đuôi này sẽ đuợc copy ra thư mục kết quả. (vd: `["out", "inp", "txt"]`)
- `testcase.output`: kí tự `*` cần được sử dùng để sau này thay thế bằng stt test. (vd: `./result/TEST*`)
- `language.[name]`: thiết lập ngôn ngữ. Xem phần duới.

### Thiết lập ngôn ngữ

Để thiết lập một ngôn ngữ, thêm một key có tên bất kì với các property sau:

- `name`: tên ngôn ngữ
- `extensions`: mảng chứa các đuôi file để nhận diện ngôn ngữ này
- `compile`: (chỉ nếu ngôn ngữ này cần đuợc biên dịch) viết array lệnh để biên dịch code. Bắt buộc bao gồm `$SOURCE` (sẽ được thay thế bằng vị trí source) và `$OUTPUT` (sẽ được thay thế bằng vị trí file binary sau khi biên dịch). Những ngôn ngữ như Python sẽ không có buớc này.
- `run`: lệnh để chạy chuơng trình. Lưu ý bao gồm `$OUTPUT` (sẽ được thay thế bằng vị trí file chạy)

Xem file thiết lập mẫu [tcgen.yml](./gen/tcgen.yml) để biết cách thiết lập.

## Viết code tạo case và giải case

Xem [ví dụ](./example/)

### Viết code tạo case

Code tạo case cần viết ra file input.

Code sẽ đuợc gọi với stt lần chạy ở argument đầu tiên (argument sau tên gọi), bắt đầu từ 1. Giá trị này có thể dùng cho các mục đích khác nhau.

Ví dụ sử dụng stt lần chạy để tạo test có độ khó khác nhau:

```cpp
// Tạo 5 test đầu dễ và các test sau khó
int main(int argc, char** argv)
{
  int caseNumber = atoi(argv[1]);
  if (caseNumber < 5) taoTestCaseDonGian();
  else taoTestCaseKho();
}
```

### Viết code giải case

Code giải case cần đọc file input đuợc tạo và viết ra file output.

Tuơng tự code sẽ đuợc gọi với stt lần chạy ở argument đầu tiên.

## Bản quyền

[MIT](LICENSE)
