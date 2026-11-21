package main

import (
    "testing"
    "slices"
)

func TestNoDiff(t *testing.T) {
    before := []Reservation{
        Reservation {
            Car: "カローラツーリングHV　車両番号0000",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ハイエースバン　車両番号1111",
            Departure: "トヨタモビリティサービス　成田空港店",
            Arrival: "トヨタレンタリース名古屋　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ヤリス　車両番号2222",
            Departure: "トヨタレンタリース青森　青森空港店",
            Arrival: "トヨタレンタリース宮城　返却可能店舗",
            DepartureDate: "",
        },
    }
    after := before

    addition_expected := []Reservation{}
    deletion_expected := []Reservation{}
    addition_actual, deletion_actual := find_diff(before, after)

    if !slices.Equal(addition_expected, addition_actual) {
        t.Errorf("Diffs in Addition\n[expected]: %v\n[actual]: %v\n", addition_expected, addition_actual)
    }

    if !slices.Equal(deletion_expected, deletion_actual) {
        t.Errorf("Diffs in Deletion\n[expected]: %v\n[actual]: %v\n", deletion_expected, deletion_actual)
    }
}

func TestSimpleAddition(t *testing.T) {
    before := []Reservation{
        Reservation {
            Car: "カローラツーリングHV　車両番号0000",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ヤリス　車両番号2222",
            Departure: "トヨタレンタリース青森　青森空港店",
            Arrival: "トヨタレンタリース宮城　返却可能店舗",
            DepartureDate: "",
        },
    }
    after := []Reservation{
        Reservation {
            Car: "カローラツーリングHV　車両番号0000",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ハイエースバン　車両番号1111",
            Departure: "トヨタモビリティサービス　成田空港店",
            Arrival: "トヨタレンタリース名古屋　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ヤリス　車両番号2222",
            Departure: "トヨタレンタリース青森　青森空港店",
            Arrival: "トヨタレンタリース宮城　返却可能店舗",
            DepartureDate: "",
        },
    }

    addition_expected := []Reservation{
        Reservation {
            Car: "ハイエースバン　車両番号1111",
            Departure: "トヨタモビリティサービス　成田空港店",
            Arrival: "トヨタレンタリース名古屋　返却可能店舗",
            DepartureDate: "",
        },
    }
    deletion_expected := []Reservation{}
    addition_actual, deletion_actual := find_diff(before, after)

    if !slices.Equal(addition_expected, addition_actual) {
        t.Errorf("Diffs in Addition\n[expected]: %v\n[actual]: %v\n", addition_expected, addition_actual)
    }

    if !slices.Equal(deletion_expected, deletion_actual) {
        t.Errorf("Diffs in Deletion\n[expected]: %v\n[actual]: %v\n", deletion_expected, deletion_actual)
    }
}

func TestSimpleDeletion(t *testing.T) {
    before := []Reservation{
        Reservation {
            Car: "カローラツーリングHV　車両番号0000",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ハイエースバン　車両番号1111",
            Departure: "トヨタモビリティサービス　成田空港店",
            Arrival: "トヨタレンタリース名古屋　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ヤリス　車両番号2222",
            Departure: "トヨタレンタリース青森　青森空港店",
            Arrival: "トヨタレンタリース宮城　返却可能店舗",
            DepartureDate: "",
        },
    }
    after := []Reservation{
        Reservation {
            Car: "ヤリス　車両番号2222",
            Departure: "トヨタレンタリース青森　青森空港店",
            Arrival: "トヨタレンタリース宮城　返却可能店舗",
            DepartureDate: "",
        },
    }

    addition_expected := []Reservation{}
    deletion_expected := []Reservation{
        Reservation {
            Car: "カローラツーリングHV　車両番号0000",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "ハイエースバン　車両番号1111",
            Departure: "トヨタモビリティサービス　成田空港店",
            Arrival: "トヨタレンタリース名古屋　返却可能店舗",
            DepartureDate: "",
        },
    }
    addition_actual, deletion_actual := find_diff(before, after)

    if !slices.Equal(addition_expected, addition_actual) {
        t.Errorf("Diffs in Addition\n[expected]: %v\n[actual]: %v\n", addition_expected, addition_actual)
    }

    if !slices.Equal(deletion_expected, deletion_actual) {
        t.Errorf("Diffs in Deletion\n[expected]: %v\n[actual]: %v\n", deletion_expected, deletion_actual)
    }
}

func TestWhenCompletelySameReservationAdded(t *testing.T) {
    before := []Reservation{
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
    }
    after := []Reservation{
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
    }

    addition_expected := []Reservation{
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
    }
    deletion_expected := []Reservation{}
    addition_actual, deletion_actual := find_diff(before, after)

    if !slices.Equal(addition_expected, addition_actual) {
        t.Errorf("Diffs in Addition\n[expected]: %v\n[actual]: %v\n", addition_expected, addition_actual)
    }

    if !slices.Equal(deletion_expected, deletion_actual) {
        t.Errorf("Diffs in Deletion\n[expected]: %v\n[actual]: %v\n", deletion_expected, deletion_actual)
    }
}

func TestWhenCompletelySameReservationDeleted(t *testing.T) {
    before := []Reservation{
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
    }
    after := []Reservation{
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
    }

    addition_expected := []Reservation{}
    deletion_expected := []Reservation{
        Reservation {
            Car: "カローラツーリングHV",
            Departure: "トヨタモビリティサービス　品川高輪口店",
            Arrival: "トヨタレンタリース青森　返却可能店舗",
            DepartureDate: "",
        },
    }
    addition_actual, deletion_actual := find_diff(before, after)

    if !slices.Equal(addition_expected, addition_actual) {
        t.Errorf("Diffs in Addition\n[expected]: %v\n[actual]: %v\n", addition_expected, addition_actual)
    }

    if !slices.Equal(deletion_expected, deletion_actual) {
        t.Errorf("Diffs in Deletion\n[expected]: %v\n[actual]: %v\n", deletion_expected, deletion_actual)
    }
}
