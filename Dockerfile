# ʹ��golang��alpine�汾��Ϊ��������
FROM golang:alpine
# ���ù���Ŀ¼
WORKDIR /app
# ����ǰĿ¼�����ݸ��Ƶ������е�/appĿ¼
COPY . .
# ����Goģ��֧�֣�������Go����
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io/zh/"
# ����GoӦ�ó���
RUN go build -o myapp main.go
# ��¶9988�˿�
EXPOSE 9988
# ������������ʱ���е�����
ENTRYPOINT ["./myapp"]