from user import ProxyService, RealService


def main():
    real = RealService()
    proxy = ProxyService(real)

    print(proxy.GetUser(10))
    print(proxy.GetUser(1))
    print(proxy.GetUser(10))
    print(proxy.GetUser(1))


if __name__ == "__main__":
    main()
