import threading

from config import Config


def access_config(thread_id):
    cfg = Config()
    print(
        f"From {thread_id} -> AppName: {cfg.AppName}, Version: {cfg.Version}, Address: {id(cfg)}"
    )


def main():
    threads = []
    for i in range(10):
        t = threading.Thread(target=access_config, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()


if __name__ == "__main__":
    main()
