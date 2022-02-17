# judgen

Tạo test case cho câu hỏi lập trình.

## Hướng dẫn

Tạo 2 file code cho các mục đích:

- **Tạo case**: Khi chạy, code này sẽ viết ra file input.
- **Giải case**: Khi chạy, code này sẽ đọc file input đuợc tạo trên và viết ra file output

Các file input và output phải có đuôi đuợc liệt kê trong thiết lập `testcase.extensions`.

Chạy chuơng trình. Sau đó nhập địa điểm 2 file trên và số lần chạy `n`.

Sau khi chạy `n` lần, chuơng trình sẽ copy hết các file input, output cho vào các thứ mục có tên dạng "TEST`i`" (trong đó `i` là số thứ tự test).

## File thiết lập

File thiết lập cần đuợc đặt trong cùng thư mục với chuơng trình `judgen`. Xem file thiết lập mẫu: [judgen.yml](./judgen.yml).

- `testcase.extensions`: file có các đuôi này sẽ đuợc copy ra thư mục kết quả.
- `output.dir`: thư mục kết quả test case
- `language.[name]`: thiết lập ngôn ngữ. Xem phần duới.

### Thiết lập ngôn ngữ

Để thiết lập một ngôn ngữ, thêm một key có tên bất kì (nên là ký tự thuờng) với các property sau:

- `name`: tên ngôn ngữ
- `extensions`: mảng chứa các đuôi file để nhận diện ngôn ngữ này
- `compile`: (chỉ nếu ngôn ngữ này cần đuợc biên dịch) viết array lệnh để biên dịch code. Bắt buộc bao gồm `SOURCE` sẽ được thay thế bằng vị trí source và `OUTPUT` là vị trí file binary sau khi biên dịch. Những ngôn ngữ như Python sẽ không có buớc này.
- `run`: lệnh để chạy chuơng trình. Lưu ý bao gồm `OUTPUT` là file sẽ đuợc chạy

Xem file thiết lập mẫu [judgen.yml](./judgen.yml) để biết cách thiết lập.

## Viết code tạo case và giải case

Xem [ví dụ](./example/)

### Viết code tạo case

Code tạo case cần viết ra file input.

Code sẽ đuợc gọi với stt lần chạy ở argument đầu tiên (argument sau tên gọi), bắt đầu từ 1. Giá trị này có thể dùng cho các mục đích khác nhau (vd: để tạo test case tuỳ theo độ khó).

Ví dụ chương trình có tổng cộng 10 case và bạn muốn 5 case dễ và 5 case khó.

```cpp
int main(int argc, char** argv)
{
  int caseNumber = atoi(argv[1]);
  if (caseNumber < 5) generateEasyTestCase();
  else generateDifficultTestCase()
}
```

### Viết code giải case

Code giải case cần đọc file input đuợc tạo và viết ra file output.

Tuơng tự code sẽ đuợc gọi với stt lần chạy ở argument đầu tiên.

## Bản quyền

[MIT](LICENSE)
