package router

import (
	"github.com/labstack/echo"
	"github.com/narozhniyy/test/internal/models"
	"github.com/narozhniyy/test/internal/resources"
	"net/http"
	"strconv"
)

// Assign new table to guests
func assignTable(c echo.Context) error {
	table := new(models.Table)
	if err := c.Bind(table); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, wrong request type/format.",
		})
	}

	lt, err := resources.GetDocument(table.Table)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.Error{
			Message: "Sorry, something went wrong.",
		})
	}
	if lt != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, this table is not available.",
		})
	}

	_, err = resources.InsertDocument(table)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, table)
}

// Make an order for each seat in table
func makeOrder(c echo.Context) error {
	table := c.Param("table")
	t, err := strconv.ParseInt(table, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, wrong request parameter.",
		})
	}

	existTable, err := resources.GetDocument(t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.Error{
			Message: "Sorry, something went wrong.",
		})
	}
	if existTable == nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, impossible to add an order for table without guests.",
		})
	}

	to := new(models.TableOrder)
	if err := c.Bind(to); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, wrong request type/format.",
		})
	}

	to.ProcessSharedOrder()
	for i, guest := range existTable.Guests {
		to.AddOrderToGuest(&guest)
		existTable.Guests[i] = guest
	}

	_, err = resources.UpdateDocument(existTable, t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.Error{
			Message: "Sorry, something went wrong.",
		})
	}

	return c.JSON(http.StatusOK, existTable)
}

// Calculate bill for provided guests in table
func getBills(c echo.Context) error {
	tn := c.Param("table")
	t, err := strconv.ParseInt(tn, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, wrong request parameter.",
		})
	}

	table, err := resources.GetDocument(t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.Error{
			Message: "Sorry, something went wrong.",
		})
	}
	if table == nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, impossible to build a bill for table without guests.",
		})
	}

	brs := new(models.BillsRequest)
	if err := c.Bind(brs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Error{
			Message: "Sorry, wrong request type/format.",
		})
	}

	bills := table.ProcessBills(brs)

	return c.JSON(http.StatusOK, bills)
}
