# judgen

Tạo test case cho câu hỏi lập trình.

## Hướng dẫn

Tạo 2 file code cho các mục đích:

- **Tạo case**: Code này sẽ chạy và viết ra file input.
- **Giải case**: Code này sẽ đọc file input đuợc tạo trên và viết ra file output

Các file input và output phải có đuôi đuợc liệt kê trong file thiết lập.

Chạy chuơng trình. Sau đó nhập địa điểm 2 file trên và số lần chạy `n`.

Sau khi chạy `n` lần, chuơng trình sẽ copy hết các file input, output cho vào các thứ mục có tên dạng TEST`ith` (trong đó `ith` là số thứ tự test).

## File thiết lập

File thiết lập cần đuợc đặt trong cùng thư mục với chuơng trình `judgen`. Xem file thiết lập mẫu: [judgen.yml](./judgen.yml).

- `testcase.extensions`: file có các đuôi này sẽ đuợc copy ra thư mục kết quả.
- `output.dir`: thư mục kết quả
- `language.[name]`: thiết lập ngôn ngữ. Xem phần duới.

### Thiết lập ngôn ngữ

Để thiết lập một ngôn ngữ, thêm một key có tên bất kì (nên là ký tự thuờng) với các property sau:

- `name`: tên ngôn ngữ
- `extensions`: mảng chứa các đuôi file của ngôn ngữ này
- `compile`: nếu ngôn ngữ này cần đuợc biên dịch, viết array lệnh để biên dịch code. Lưu ý bao gồm `SOURCE` là file code và `OUTPUT` là file chuơng trình đuợc biên dịch. Nhưng ngôn ngữ như Python sẽ không có buớc này.
- `run`: lệnh để chạy chuơng trình. Lưu ý bao gồm `OUTPUT` là file sẽ đuợc chạy

Xem file thiết lập mẫu: [judgen.yml](./judgen.yml) để biết cách thiết lập.

## Viết code tạo case

Code tạo case cần viết ra file input.

Chuơng trình sẽ đuợc gọi với stt lần chạy ở argument đầu tiên (argument sau tên chuơng trình), bắt đầu từ 0. Giá trị này có thể dùng để tạo test case cho các mục đích (vd: tuỳ theo độ khó).

Ví dụ chương trình có tổng cộng 10 case và bạn muốn 5 case dễ và 5 case khó.

```cpp
int main(int argc, char** argv)
{
  int caseNumber = atoi(argv[1]);
  if (caseNumber < 5) generateEasyTestCase();
  else generateDifficultTestCase()
}
```

## Viết code tạo case

Code tạo case cần đọc file input đuợc tạo và viết ra file output.

Tuơng tự chuơng trình sẽ đuợc gọi với stt lần chạy ở argument đầu tiên.

## Bản quyền

[MIT](LICENSE)
