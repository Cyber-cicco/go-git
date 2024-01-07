package entites

type Head struct {
    Ref string `json:"ref"`
}

func GetBaseHead() Head {
    return Head{
        Ref: "refs/head/master",
    }
}
