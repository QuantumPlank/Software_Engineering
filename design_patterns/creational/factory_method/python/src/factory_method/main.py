from kitchen import Kitchen

if __name__ == "__main__":
    k = Kitchen()
    meal1 = k.create_meal("pizza")
    print(meal1.serve())
    meal2 = k.create_meal("hamburger")
    print(meal2.serve())
